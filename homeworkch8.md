# 1、Service

```
apiVersion: v1
kind: Service
metadata:
  name: httpserver
  labels:
    app: httpserver
spec:
  ports:
    - port: 8080
      targetPort: 8080
      name: "http"
  selector:
    run: httpserver
---
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    run: httpserver
  name: httpserver
spec:
  replicas: 2
  selector:
    matchLabels:
      run: httpserver
  template:
    metadata:
      annotations:
        prometheus.io/scrape: "true"
        prometheus.io/port: "8080"
      labels:
        run: httpserver
    spec:
      imagePullSecrets:
        - name: harbor-registry-secret
      containers:
        - image: mocking/httpserver:1.1
          name: httpserver
          ports:
            - containerPort: 8080
          env:
            - name: VERSION
              valueFrom:
                configMapKeyRef:
                  key: version
                  name: httpserver
          startupProbe:
            httpGet:
              port: 8080
              path: /healthz
            initialDelaySeconds: 5
            periodSeconds: 5
            timeoutSeconds: 1
          readinessProbe:
            httpGet:
              port: 8080
              path: /healthz
            periodSeconds: 5
            timeoutSeconds: 1
          livenessProbe:
            httpGet:
              port: 8080
              path: /healthz
            periodSeconds: 5
            timeoutSeconds: 1
          resources:
            limits:
              memory: 256Mi
              cpu: 200m
            requests:
              memory: 128Mi
              cpu: 100m
```

# 2、Ingress

```
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: gateway
  annotations:
    kubernetes.io/ingress.class: "nginx"
spec:
  tls:
    - hosts:
        - cncamp.com
      secretName: cncamp-tls
  rules:
    - host: cncamp.com
      http:
        paths:
          - path: "/"
            pathType: Prefix
            backend:
              service:
                name: httpserver
                port:
                  number: 8080
```

