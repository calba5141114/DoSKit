package server

import (
	"log"
	"github.com/gorilla/mux"
	"net/http"
)


func attackHandler(w http.ResponseWriter, r *http.Request){

}	

func indexHandler(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/html; charset=UTF-8")
}


func main(){
	r := mux.NewRouter()
	r.HandleFunc("/DosAttack", attackHandler ).Methods("POST")
	r.HandleFunc("/", indexHandler).Methods("GET")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}