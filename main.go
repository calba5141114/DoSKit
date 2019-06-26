package main

import (
	// "net/http"
	"flag"
  "fmt"
  "os"
)

func main() {
  v := flag.String("target", "", "target to run DoS Attack on")
  
  if *v == "" {
    flag.PrintDefaults()

  }
	flag.Parse()
	fmt.Println(flag.Args())
}
