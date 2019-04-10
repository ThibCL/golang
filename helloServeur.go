//Package main ...
package main

import (
	"net/http"

	"github.com/ThibCL/gotest/serveur"
	"github.com/ThibCL/gotest/store"

	"github.com/gorilla/mux"
)

var s store.Store

func main() {
	serveur.InitializeStore()
	r := mux.NewRouter()
	r.HandleFunc("/hello", serveur.SayHello).Methods("GET")
	r.HandleFunc("/hello", serveur.AddHello).Methods("POST")
	r.HandleFunc("/hello", serveur.DeleteHello).Methods("DELETE")
	http.ListenAndServe(":9000", r)
}
