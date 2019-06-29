package server

import (
	"github.com/gorilla/mux"
	"net/http"
)


func attackHandler(w http.ResponseWriter, r *http.Request){

}	

func main(){
	r := mux.NewRouter()
	r.HandleFunc("/DosAttack", attackHandler ).Methods("POST")
}