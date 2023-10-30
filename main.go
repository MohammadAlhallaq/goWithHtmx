package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	services "goWithHtmx/model/service"
	"log"
	"net/http"
)

func sendItems(w http.ResponseWriter) {
	is, _ := services.NewItemService(services.WithMysqlItem("root:@tcp(127.0.0.1:3306)/hmshop"))
	fmt.Println(is.GetAllItems())
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//items, err := is.Item.GetAll()
	//fmt.Println(items)

	//users, _ := model.GetAllItems()
	//fmt.Println(users)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//err = template.Must(template.ParseFiles("src/public/index.html")).ExecuteTemplate(w, "items", items)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
}

func index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sendItems(w)
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
