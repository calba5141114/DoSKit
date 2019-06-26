package main

import (
	// "net/http"
	"flag"
	"fmt"
)

func main() {
	v := flag.String("target", "", "target to run DoS Attack on")
	fmt.Println(v)
	flag.Parse()
	fmt.Println(flag.Args())
}
