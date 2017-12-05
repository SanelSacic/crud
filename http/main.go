package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/gocode/crud/http/web"
	"github.com/gocode/crud/http/web/models"
)

func main() {

	db, err := models.NewMysql("root:sanel@/bookshell")
	if err != nil {
		log.Println(err)
	}

	cru := &web.Crud{db}

	r := mux.NewRouter()

	// --------- >
	r.HandleFunc("/index", cru.ServeIndexPage).Methods("GET")
	// --------- >
	r.HandleFunc("/authors", cru.ListAuthors).Methods("GET")
	r.HandleFunc("/edit-author/{id:[0-9]+}", cru.EditAuthor).Methods("GET")
	r.HandleFunc("/update-author", cru.UpdateAuthor).Methods("POST")
	r.HandleFunc("/delete-author/{id:[0-9]+}", cru.DeleteAuthor).Methods("GET")
	r.HandleFunc("/create-author", cru.CreateAuthor).Methods("POST")
	// --------- >
	r.HandleFunc("/books", cru.ListBooks).Methods("GET")
	r.HandleFunc("/edit-book/{id:[0-9]+}", cru.EditBook).Methods("GET")
	r.HandleFunc("/update-book", cru.UpdateBook).Methods("POST")
	r.HandleFunc("/delete-book/{id:[0-9]+}", cru.DeleteBook).Methods("GET")
	r.HandleFunc("/create-book", cru.CreateBook).Methods("POST")
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Println("Listen and Serve ...", err)
	}
}
