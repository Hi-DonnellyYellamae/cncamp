package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func index(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("<h1>Welcome to cloudnative！</h1>"))

	os.Setenv("VERSION", "0.1")
	version := os.Getenv("VERSION")
	w.Header().Set("VERSION", version)
	fmt.Printf("os version: %s\n", version)

	for k, v := range r.Header {
		for _, vv := range v {
			fmt.Printf("Header Key: %s, Header value: %s\n", k, v)
			w.Header().Set(k, vv)
		}
	}

	clientIP := getCurrentIP(r)
	log.Printf("Response code: %d", http.StatusAccepted)
	log.Printf("clientIP:%s", clientIP)
}

func healthz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "working")
}

func getCurrentIP(r *http.Request) string {
	ip := r.Header.Get("X-REAL-IP")
	if ip == "" {
		//ip:port
		ip = strings.Split(r.RemoteAddr, ":")[0]
	}
	return ip
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", index)
	mux.HandleFunc("/healthz", healthz)
	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		log.Fatalf("Start httpserver failed, error: %v\n", err.Error())
	}
}
