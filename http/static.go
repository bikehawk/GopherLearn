package main

import (
	"fmt"
	"log"
	"net/http"
)

func staticHandler(w http.ResponseWriter, r *http.Request) {
	//http.Handle(r.Host+"/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("/var/testfiles/static"))))
	fs := http.StripPrefix("/static/", http.FileServer(http.Dir("/var/testfiles/static")))
	fmt.Println(fs)
	fs.ServeHTTP(w, r)
}

func main() {

	http.HandleFunc("/static/", staticHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
