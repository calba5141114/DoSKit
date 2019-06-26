package main 
import (
  // "net/http"
  "flag"
  "fmt"
)

func main(){
  v := flag.String("target", "", "target to run DoS Attack on")
  flag.Parse()
  fmt.Println()
}
