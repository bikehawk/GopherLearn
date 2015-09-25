package main

import (
	"fmt"
	"github.com/dvsekhvalnov/jose2go/keys/rsa"
	"io/ioutil"
)

func main() {

	keyBytes, _ := ioutil.ReadFile("private.key")

	privateKey, err := Rsa.ReadPrivate(keyBytes)

	if err != nil {
		panic("invalid key format")
	}

	fmt.Printf("privateKey = %v\n", privateKey)
}
