package main

import (
	"fmt"
	"github.com/dvsekhvalnov/jose2go"
	"github.com/dvsekhvalnov/jose2go/keys/rsa"
	"io/ioutil"
	"os"
)

func main() {

	token := os.Args[1]

	keyBytes, err := ioutil.ReadFile("private.key")

	if err != nil {
		panic("invalid key file")
	}

	privateKey, e := Rsa.ReadPrivate(keyBytes)

	if e != nil {
		panic("invalid key format")
	}

	payload, headers, err := jose.Decode(token, privateKey)

	if err == nil {
		//go use payload
		fmt.Printf("\npayload = %v\n", payload)

		//and/or use headers
		fmt.Printf("\nheaders = %v\n", headers)
	} else {
		fmt.Printf("Could not verify token: %v\n", token)
	}
}
