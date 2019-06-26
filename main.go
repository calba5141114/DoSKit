package main

import (
	// "net/http"
	"flag"
	"fmt"
	"os"
)

func main() {

	if *v == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	flag.Parse()
	fmt.Println(flag.Args())
}
