package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
)

var hashKey = []byte("12345678901234567890123456789012")
var blockKey = []byte("1234567890123456")
var s = securecookie.New(hashKey, blockKey)

func SetCookieHandler(w http.ResponseWriter, r *http.Request) {
	value := map[string]string{
		"foo": "bar",
	}
	if encoded, err := s.Encode("cookie-name", value); err == nil {
		cookie := &http.Cookie{
			Name:  "cookie-name",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(w, cookie)
		fmt.Fprint(w, "Cookie set!")
	} else {
		log.Println(err)
		fmt.Fprintf(w, "Could not set cookie, returned error: %s", err)
	}
}

func ReadCookieHandler(w http.ResponseWriter, r *http.Request) {
	if cookie, err := r.Cookie("cookie-name"); err == nil {
		value := make(map[string]string)
		if err = s.Decode("cookie-name", cookie.Value, &value); err == nil {
			fmt.Fprintf(w, "The value of foo is %q", value["foo"])
		}
	} else {
		log.Println(err)
		fmt.Fprintf(w, "Cookie validation failed, returned error: %s", err)
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/getcookie", ReadCookieHandler).Methods("GET")
	router.HandleFunc("/setcookie", SetCookieHandler).Methods("GET")

	err := http.ListenAndServe(":8089", router)
	if err != nil {
		log.Fatalf("Could not start the server, returned error: %s\n", err)
	}
}
