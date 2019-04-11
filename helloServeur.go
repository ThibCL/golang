//Package main ...
package main

import (
	"net/http"

	"github.com/ThibCL/gotest/serveur"

	"github.com/ThibCL/gotest/store"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	str := store.NewStore()
	helloService := serveur.NewHelloService(&str)
	helloService.Register(r)
	http.ListenAndServe(":9000", r)
}
