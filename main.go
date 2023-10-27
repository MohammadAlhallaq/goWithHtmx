package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"goWithHtmx/model"
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	template.Must(template.ParseFiles("./index.html")).Execute(w, nil)
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func main() {

	model.Setup()
	//router := httprouter.New()
	//
	//router.GET("/", Index)
	//router.GET("/hello/:name", Hello)
	//
	//log.Fatal(http.ListenAndServe(":8080", router))
}
