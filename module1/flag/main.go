package main

import (
	"flag"
	"fmt"
)

var mode = flag.String("mode", "", "fast模式")

func main() {
	flag.Parse()
	fmt.Println(*mode)
}
