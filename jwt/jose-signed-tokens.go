package main

import (
	"fmt"
	"github.com/dvsekhvalnov/jose2go"
)

func main() {

	payload := `{"hello": "world"}`

	key := []byte{97, 48, 97, 50, 97, 98, 100, 56, 45, 54, 49, 54, 50, 45, 52, 49, 99, 51, 45, 56, 51, 100, 54, 45, 49, 99, 102, 53, 53, 57, 98, 52, 54, 97, 102, 99}

	token, err := jose.Sign(payload, jose.HS256, key)

	if err == nil {
		//go use token
		fmt.Printf("\nHS256 = %v\n", token)
	}
}
