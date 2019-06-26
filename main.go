package main 
import (
  // "net/http"
  "flag"
  "fmt"
)

func main(){
  flag.String("target", "", "target to run DoS Attack on")
  flag.Parse()
  fmt.Println(flag.Args())
}
