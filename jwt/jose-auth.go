package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/dvsekhvalnov/jose2go"
	"github.com/dvsekhvalnov/jose2go/keys/rsa"
	"io/ioutil"
	"log"
)

var username string
var password string
var tokenToVerify string

var users = map[string]string{
	"admin": "adminpass",
	"user1": "user1pass",
	"user2": "user2pass",
}

func init() {
	const (
		defaultUser  = "anonymous"
		defaultPass  = "null"
		defaultToken = "null"
	)
	flag.StringVar(&username, "user", defaultUser, "Your username.")
	flag.StringVar(&username, "u", defaultUser, "Your username. (shorthand)")
	flag.StringVar(&password, "pass", defaultPass, "Your password.")
	flag.StringVar(&password, "p", defaultPass, "Your password. (shorthand)")
	flag.StringVar(&tokenToVerify, "verify", defaultToken, "The token to verify.")
	flag.StringVar(&tokenToVerify, "v", defaultToken, "The token to verify. (shorthand)")
}

func verifyUser(username, password string) bool {
	if _, ok := users[username]; ok {
		if password == users[username] {
			return true
		}
	}
	return false
}

func getToken() (string, error) {
	payload := `{"site": "jose-auth.go"}`

	keyBytes, err := ioutil.ReadFile("public.cer")
	if err != nil {
		return "", errors.New("Oh noes! Seems the certificate is invalid! User cannot has access!")
	}
	publicKey, err := Rsa.ReadPublic(keyBytes)
	if err != nil {
		return "", errors.New("Oh noes! Could not read the public key! User cannot has access!")
	}
	token, err := jose.Encrypt(payload, jose.RSA_OAEP, jose.A256GCM, publicKey)
	if err != nil {
		return "", errors.New("Oh noes! Could not generate the token! User cannot has access!")
	}
	return token, nil
}

func verifyToken(token string) (string, map[string]interface{}, error) {
	keyBytes, err := ioutil.ReadFile("private.key")
	if err != nil {
		return "", nil, err
	}

	privateKey, err := Rsa.ReadPrivate(keyBytes)
	if err != nil {
		return "", nil, err
	}

	payload, headers, err := jose.Decode(token, privateKey)
	if err == nil {
		return payload, headers, nil
	}
	return "", nil, err
}

func main() {
	flag.Parse()

	if username == "anonymous" && tokenToVerify == "null" {
		log.Println("You are anonymous! You cannot has token!")
	} else {
		if username == "anonymous" {
			payload, headers, err := verifyToken(tokenToVerify)
			if err != nil {
				log.Println("Could not verify token! You cannot has access!")
				log.Printf("Error: %s\n", err)
			} else {
				fmt.Println("The token is valid! You can has access!")
				fmt.Printf("\npayload: %v\n", payload)
				fmt.Printf("\nheaders: %v\n", headers)
			}
		} else {
			if verifyUser(username, password) {
				token, err := getToken()
				if err != nil {
					log.Println(err)
				} else {
					fmt.Println("Username match! You can has access!")
					fmt.Printf("\nYour token is...\n\n%s\n", token)
				}
			} else {
				fmt.Println("Invalid username or password! You cannot has access!")
			}
		}
	}
}
