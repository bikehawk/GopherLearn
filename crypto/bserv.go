package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

func handleLogin(w http.ResponseWriter, r *http.Request) {

	dbURL = "http://localhost:8082"

	var data struct {
		Password string `json:"password"`
		PassHash string `json:"passhash"`
	}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	json.Unmarshal(b, &data)

	res, err := http.Post(fmt.Sprint(dbURL, "/userbyname"), "application/json", bytes.NewBuffer(b))
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}

	var u struct {
		ID       int    `json:"id"`
		UserName string `json:"username"`
		PassHash string `json:"passhash"`
		Role     int    `json:"role_id"`
	}

	// Comparing the password with the hash
	err = bcrypt.CompareHashAndPassword([]byte(data.PassHash), []byte(data.Password))
	if err != nil {
		fmt.Printf("Password %s does not match hash %s\n", data.Password, data.PassHash)
	} else {
		fmt.Printf("Password %s matches hash %s\n", data.Password, data.PassHash)
	}
}

func handleNew(w http.ResponseWriter, r *http.Request) {

	var data struct {
		Password string `json:"password"`
		PassHash string `json:"passhash"`
	}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	json.Unmarshal(b, &data)

	// Hashing with a default cost of 10
	passHash, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	data.PassHash = string(passHash)
	fmt.Printf("The hash for password %s is %s\n", data.Password, data.PassHash)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/login", handleLogin).Methods("POST")
	router.HandleFunc("/new", handleNew).Methods("POST")

	err := http.ListenAndServe(":8089", router)
	if err != nil {
		log.Fatalf("Could not start the server, returned error: %s\n", err)
	}
}
