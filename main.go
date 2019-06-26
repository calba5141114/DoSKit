package main

import (
	// "net/http"
	"flag"
	"fmt"
	"os"
)

var v = flag.String("target", "", "target to run DoS Attack on")

func main() {
  flag.Parse()
	if *v == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
  fmt.Println(*v)
  go func(){

  }()
}
