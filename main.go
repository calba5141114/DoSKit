package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

var v = flag.String("target", "", "target to run DoS Attack on")

func main() {
  responseChannel := make(chan *Response)
	flag.Parse()
	if *v == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	fmt.Println(*v)
	go func() {
		resp, err := http.Get(*v)
		if err != nil {
			panic(err)
		}
		fmt.Println(resp)
	}()
}
