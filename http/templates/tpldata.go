package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

func handleMain(w http.ResponseWriter, r *http.Request) {

	data := make(map[string]interface{})

	data["string"] = "my test string"
	data["int"] = 7

	t, err := template.ParseFiles("./layout.html", "./content.html")
	if err != nil {
		log.Fatalln(err)
	} else {
		t.ExecuteTemplate(w, "layout", data)
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.Methods("GET").Path("/").Name("main").HandlerFunc(handleMain)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalln(err)
	}
}
