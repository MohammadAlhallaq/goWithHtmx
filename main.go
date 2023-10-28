package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	template.Must(template.ParseFiles("src/public/index.html")).Execute(w, nil)
}

func create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func getItems(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func main() {
	setupAndRun()
}

func setupAndRun() {

	router := httprouter.New()

	router.ServeFiles("/src/*filepath", http.Dir("src"))

	router.GET("/items", index)
	router.PUT("/items/:id", update)
	router.DELETE("/items/:id", delete)
	router.POST("/items/create", create)
	log.Fatal(http.ListenAndServe(":8080", router))
}
