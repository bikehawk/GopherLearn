package main

import (
	"fmt"
	"log"
	"net/http"
)

type Data struct {
	Host, Path, Method, Proto, RemoteIP string
	Header                              http.Header
}

func mainHandler(w http.ResponseWriter, r *http.Request) {

	data := &Data{
		Host:     r.Host,
		Path:     r.URL.Path,
		Method:   r.Method,
		Proto:    r.Proto,
		RemoteIP: r.RemoteAddr,
		Header:   r.Header,
	}

	fmt.Fprintf(w, "%#v", data)
}

func main() {
	http.HandleFunc("/", mainHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
