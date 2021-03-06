package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Println("Listen and Serve ...", err)
	}
}
