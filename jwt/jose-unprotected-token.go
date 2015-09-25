package main

import (
	"fmt"
	"github.com/dvsekhvalnov/jose2go"
)

func main() {

	payload := `{"hello": "world"}`

	token, err := jose.Sign(payload, jose.NONE, nil)

	if err == nil {
		//go use token
		fmt.Printf("\nPlaintext = %v\n", token)
	}
}
