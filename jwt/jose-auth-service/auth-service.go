package main

import (
	"./jjwt"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/matryer/respond.v1"
	"log"
	"net/http"
)

type User struct {
	Name         string `json:"name"`
	PasswordHash string `json:"-"`
	Role         string `json:"role"`
}

var usermap = map[string]User{
	"admin": &User{"Big Kahuna", "$2a$10$mK7OTVwUTEzTqgdZqCscousngVDJsjW5hUl4Zs/ZBlwtABSG8Zm7q", "Administrator"},
	"user1": &User{"Little Piggy", "$2a$10$hz/AEA95ahPcrWGUqJFzEO1JzivSi2doSLlXsGGOwuf7iKSkmD7di", "User"},
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	respond.With(w, r, http.StatusNotFound, `{"status": "404"}`)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.ParseForm["username"][0]
	password := r.ParseForm["password"][0]
	if _, ok := usermap[username]; ok {
		err := bcrypt.CompareHashAndPassword(usermap[username].PasswordHash, password)
		if err == nil {
			// get a token
		}
	}
}

func handleVerify(w http.ResponseWriter, r *http.Request) {

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", handleIndex)
	router.HandleFunc("/login", handleLogin).Methods("POST")
	router.HandleFunc("/verify", handleVerify).Methods("POST")

	http.ListenAndServe(":8080", router)
}

func bcrypt(password string) {

}
