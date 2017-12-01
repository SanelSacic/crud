package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/gorilla/mux"

	"github.com/gocode/crud/http/web/models"
)

type Crud struct {
	db models.Cruder
}

type Pagedata struct {
	EntryA []models.Author
	EntryB []models.Book
}

var t = template.Must(template.ParseFiles("theme/index.html", "theme/authors.html", "theme/editAuthor.html", "theme/books.html", "theme/editBook.html"))

func renderTemplate(w http.ResponseWriter, temp string, p *Pagedata) {
	if err := t.ExecuteTemplate(w, temp+".html", &p); err != nil {
		log.Printf("Error rendering template [%s] : %s", temp, err)
	}
}

func main() {

	db, err := models.NewMysql("root:sanel@/bookshell")
	if err != nil {
		log.Println(err)
	}

	cru := &Crud{db}

	r := mux.NewRouter()
	// --------- >
	r.HandleFunc("/index", cru.serveIndexPage).Methods("GET")
	// --------- >
	r.HandleFunc("/authors", cru.serveAuthorPage).Methods("GET")
	r.HandleFunc("/edit-author/{id:[0-9]+}", cru.editAuthor).Methods("GET")
	r.HandleFunc("/update-author", cru.updateAuthor).Methods("POST")
	r.HandleFunc("/delete-author/{id:[0-9]+}", cru.deleteAuthor).Methods("GET")
	r.HandleFunc("/create-author", cru.createAuthor).Methods("POST")
	// --------- >
	r.HandleFunc("/books", cru.serveBookPage).Methods("GET")
	r.HandleFunc("/edit-book/{id:[0-9]+}", cru.editBook).Methods("GET")
	r.HandleFunc("/update-book", cru.updateBook).Methods("POST")
	r.HandleFunc("/delete-book/{id:[0-9]+}", cru.DeleteBook).Methods("GET")
	r.HandleFunc("/create-book", cru.createBook).Methods("POST")
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Println("Listen and Serve ...", err)
	}
}

func (c *Crud) serveIndexPage(w http.ResponseWriter, r *http.Request) {
	p := Pagedata{}
	w.Header().Set("Content-Type", "text/html")
	renderTemplate(w, "index", &p)
}

// -------------------------------------------------------------------->
// Authors

func (c *Crud) serveAuthorPage(w http.ResponseWriter, r *http.Request) {

	data, err := c.db.ListAuthors()
	if err != nil {
		log.Println(err)
	}
	//	log.Println(string(data))

	var authors []models.Author

	if err := json.Unmarshal(data, &authors); err != nil {
		log.Println(err)
	}

	p := Pagedata{
		EntryA: authors,
	}
	w.Header().Set("Content-Type", "text/html")
	renderTemplate(w, "authors", &p)
}

func (c *Crud) editAuthor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	data, err := c.db.RetrieveAuthor(vars["id"])
	if err != nil {
		log.Println(err)
	}

	var a []models.Author

	if err = json.Unmarshal(data, &a); err != nil {
		log.Println(err)
	}

	p := Pagedata{
		EntryA: a,
	}

	w.Header().Set("Content-Type", "text/html")
	renderTemplate(w, "editAuthor", &p)
}

func (c *Crud) updateAuthor(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println("Error while parsing form :", err)
	}

	aut := models.Author{
		ID:          r.FormValue("id"),
		Name:        r.FormValue("name"),
		LastName:    r.FormValue("lastname"),
		Description: r.FormValue("description"),
		Birth:       r.FormValue("birth"),
	}

	if aut.IsValid() {
		if _, err := c.db.UpdateAuthor(&aut); err != nil {
			log.Println(err)
		}
		log.Println("Succ")
	}

	http.Redirect(w, r, "/authors", http.StatusFound)
}

func (c *Crud) deleteAuthor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if err := c.db.DeleteAuthor(vars["id"]); err != nil {
		log.Println(err)
	}

	http.Redirect(w, r, "/authors", http.StatusFound)
}

func (c *Crud) createAuthor(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm; err != nil {
		log.Println("Error while parsing form :", err)
	}

	aut := models.Author{
		Name:        r.FormValue("name"),
		LastName:    r.FormValue("lastname"),
		Description: r.FormValue("description"),
		Birth:       r.FormValue("birth"),
	}

	log.Printf("A : [%+v]", aut)

	if aut.IsValid() {
		if _, err := c.db.CreateAuthor(&aut); err != nil {
			log.Println(err)
		}
		log.Println("Succ")
	}

	http.Redirect(w, r, "/authors", http.StatusFound)
}

// ---------------------------------------------------->
// Books

func (c *Crud) serveBookPage(w http.ResponseWriter, r *http.Request) {
	data, err := c.db.ListBooks()
	if err != nil {
		log.Println(err)
	}

	var books []models.Book

	if err = json.Unmarshal(data, &books); err != nil {
		log.Println("Error while unmarshaling data :", err)
	}

	p := Pagedata{
		EntryB: books,
	}

	w.Header().Set("Content-Type", "text/html")
	renderTemplate(w, "books", &p)
}

func (c *Crud) editBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	data, err := c.db.RetrieveBook(vars["id"])
	if err != nil {
		log.Println(err)
	}

	var books []models.Book

	if err = json.Unmarshal(data, &books); err != nil {
		log.Println("Error while unmarshaling data : ", err)
	}

	p := Pagedata{
		EntryB: books,
	}

	w.Header().Set("Content-Type", "text/html")
	renderTemplate(w, "editBook", &p)
}

func (c *Crud) updateBook(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseMultipartForm(32 << 22); err != nil {
		log.Println("Error : ParseMultipartForm :", err)
	}

	file, e, err := r.FormFile("image")
	if err != nil {
		log.Println("Error : r.FormFile :", err)
	}
	defer file.Close()

	f, err := os.OpenFile("./assets/img/"+e.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println("Error : os.OpenFile :", err)
	}
	defer f.Close()

	io.Copy(f, file)

	b := models.Book{
		ID:          r.FormValue("id"),
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
		Published:   r.FormValue("published"),
		Image:       e.Filename,
		AuthorID:    r.FormValue("author-id"),
	}

	log.Printf("%+v", b)

	if _, err = c.db.UpdateBook(&b); err != nil {
		log.Println(err)
	}

	http.Redirect(w, r, "/books", http.StatusFound)
}

func (c *Crud) DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if err := c.db.DeleteBook(vars["id"]); err != nil {
		log.Println(err)
	}

	http.Redirect(w, r, "/books", http.StatusFound)

}

func (c *Crud) createBook(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(32 << 22); err != nil {
		log.Println("Error : r.ParseMultipartForm :", err)
	}

	file, e, err := r.FormFile("image")
	if err != nil {
		log.Println("Error : r.FormFile :", err)
	}
	defer file.Close()

	f, err := os.OpenFile("./assets/img/"+e.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	io.Copy(f, file)

	book := models.Book{
		ID:          r.FormValue("id"),
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
		Published:   r.FormValue("published"),
		Image:       e.Filename,
		AuthorID:    r.FormValue("author-id"),
	}

	if _, err := c.db.CreateBook(&book); err != nil {
		log.Println(err)
	}

	http.Redirect(w, r, "/books", http.StatusFound)
}
