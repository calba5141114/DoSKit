package main

import (
	// "net/http"
	"flag"
	"fmt"
	"os"
)

var v = flag.String("target", "", "target to run DoS Attack on")

func main() {

	if *v == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	flag.Parse()
	fmt.Println(flag.Args())
}
