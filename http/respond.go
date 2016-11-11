package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gopkg.in/matryer/respond.v1"
	"log"
	"net/http"
)

type User struct {
	Role      string `json:"role"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var userlist = map[string]*User{
	"admin": &User{
		Role:      "Administrator",
		FirstName: "John",
		LastName:  "Doe",
	},
	"user1": &User{
		Role:      "User",
		FirstName: "Jane",
		LastName:  "Doe",
	},
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	respond.With(w, r, http.StatusNotFound, `{"status": "404"}`)
}

func handleUserList(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(&userlist); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

/*func handleUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user := vars["username"]
}*/

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", handleIndex)
	router.HandleFunc("/user", handleUserList)
	//router.HandleFunc("/user/{username:[a-z]+}", handleUser)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalf("Could not start the server, returned error: %s\n", err)
	}
}
