package main

import (
	"fmt"
	"github.com/dvsekhvalnov/jose2go/keys/rsa"
	"io/ioutil"
)

func main() {

	keyBytes, _ := ioutil.ReadFile("public.cer")

	publicKey, err := Rsa.ReadPublic(keyBytes)

	if err != nil {
		panic("invalid key format")
	}

	fmt.Printf("publicKey = %v\n", publicKey)
}
