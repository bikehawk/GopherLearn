package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := []byte("adminpass")

	// Hasing with a default cost of 10
	hashedPass, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(hashedPass))

	// Comparing the password with the hash
	err = bcrypt.CompareHashAndPassword(hashedPass, password)
	if err != nil {
		fmt.Println("Password does not match!")
	} else {
		fmt.Println("Password match!")
	}
}
