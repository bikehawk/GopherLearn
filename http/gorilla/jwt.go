package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"

	"github.com/dvsekhvalnov/jose2go"
	"github.com/dvsekhvalnov/jose2go/keys/rsa"
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
)

// check for cookie, if no cookie show login
// if cookie, check token validity
// if token is valid show content, else show login
// gen token upon login
// store token in cookie

//var hashKey = []byte("USanLhAWDsq^zW@uhUhsa$&!6L5Ht9GdzGRn3$L-De6_x^8Pu47kV?QGn?9@vXE+")
//var blockKey = []byte("22hYWa9?_hPmg#aRKT=nm2@-XU^yBuU?")
//var s = securecookie.New(hashKey, blockKey)
var s = securecookie.New(securecookie.GenerateRandomKey(64), securecookie.GenerateRandomKey(32))

func Authenticate(r *http.Request) (string, error) {
	value := make(map[string]string)
	if cookie, err := r.Cookie("jwt-test-cookie"); err == nil {
		if err = s.Decode("jwt-test-cookie", cookie.Value, &value); err == nil {
			payload, _, err := verifyToken(value["token"], "./private.key")
			if err != nil {
				// Invalid token!
				return "", errors.New("Invalid JWT token!")
			} else {
				// Success! We have been authenticated!
				return payload, nil
			}
		} else {
			// Invalid cookie!
			return "", errors.New("Invalid cookie!")
		}
	} else {
		// No cookie.
		return "", errors.New("No cookie.")
	}
}

func SetCookie(w http.ResponseWriter, r *http.Request, token string) {
	value := map[string]string{
		"token": token,
	}
	if encoded, err := s.Encode("jwt-test-cookie", value); err == nil {
		cookie := &http.Cookie{
			Name:  "jwt-test-cookie",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(w, cookie)
	} else {
		log.Println(err)
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	_, err := Authenticate(r)
	if err != nil {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}

		t := `<div>
			<form action="/login" method="POST">
				<input type="text" name="username">
				<input type="password" name="password">
				<input type="submit" value="Submit">
			</form>
		</div>`

		tpl := template.New("Template Switch")
		username, found := r.Form["username"]
		password := r.Form["password"]

		if found {
			if username[0] == "admin" && password[0] == "adminpass" {
				// get token and redirect to test page.
				token, err := Token("{\"role\": \"1\"}", "./public.cer")
				if err != nil {
					fmt.Println(err)
				} else {
					SetCookie(w, r, token)
					http.Redirect(w, r, "/test", 303)
				}
			}
		}
		tpl, err = tpl.Parse(t)
		if err != nil {
			log.Println(err)
		}

		err = tpl.Execute(w, nil)
	} else {
		// We are already authenticated!
		http.Redirect(w, r, "/test", 303)
	}
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	payload, err := Authenticate(r)
	if err != nil {
		// auth failed!
		http.Redirect(w, r, "/login", 303)
	} else {
		// auth success!
		fmt.Fprintf(w, "<h1>Authenticated!</h1><div>Payload: %s", payload)
	}
}

func Token(payload, pubCert string) (string, error) {
	// Read the public certificate file.
	keyBytes, err := ioutil.ReadFile(pubCert)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Could not read public certificate. Returned error: %s", err))
	}

	// Read the public key we've read from the public certificate file.
	publicKey, err := Rsa.ReadPublic(keyBytes)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Could not read public key. Returned error: %s", err))
	}

	// Generate a token using the public key.
	token, err := jose.Encrypt(payload, jose.RSA_OAEP, jose.A256GCM, publicKey)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Could not generate token. Returned error: %s", err))
	}
	return token, nil
}

func verifyToken(token, privKey string) (string, map[string]interface{}, error) {
	// Read the private key file.
	keyBytes, err := ioutil.ReadFile(privKey)
	if err != nil {
		return "", nil, errors.New(fmt.Sprintf("Could not read private key file. Returned error: %s", err))
	}

	// Read the private key we've read from the private key file.
	privateKey, err := Rsa.ReadPrivate(keyBytes)
	if err != nil {
		return "", nil, errors.New(fmt.Sprintf("Could not read private key. Returned error: %s", err))
	}

	// Decode the token using the private key.
	payload, headers, err := jose.Decode(token, privateKey)
	if err != nil {
		return "", nil, errors.New(fmt.Sprintf("Could not verify token. Returned error: %s", err))
	}
	return payload, headers, nil
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/login", LoginHandler).Methods("GET", "POST")
	router.HandleFunc("/test", TestHandler).Methods("GET")

	err := http.ListenAndServe(":8089", router)
	if err != nil {
		log.Fatalf("Could not start the server, returned error: %s\n", err)
	}
}
