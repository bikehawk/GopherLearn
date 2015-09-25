package main

import (
	"fmt"
	"github.com/dvsekhvalnov/jose2go"
	"github.com/dvsekhvalnov/jose2go/keys/rsa"
	"io/ioutil"
)

func main() {

	payload := `{"hello": "world"}`

	keyBytes, err := ioutil.ReadFile("public.cer")

	if err != nil {
		panic("invalid key file")
	}

	publicKey, e := Rsa.ReadPublic(keyBytes)

	if e != nil {
		panic("invalid key format")
	}

	//OR:
	//token,err := jose.Encrypt(payload, jose.RSA1_5, jose.A256GCM, publicKey)
	token, err := jose.Encrypt(payload, jose.RSA_OAEP, jose.A256GCM, publicKey)

	if err == nil {
		//go use token
		fmt.Printf("\ntoken = %v\n", token)
	}
}
