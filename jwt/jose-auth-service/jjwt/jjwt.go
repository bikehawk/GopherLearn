// Package jjwt uses the jose2go library to generate and verify Json Web Tokens.
package jjwt

import (
	"errors"
	"fmt"
	"github.com/dvsekhvalnov/jose2go"
	"github.com/dvsekhvalnov/jose2go/keys/rsa"
	"io/ioutil"
)

// Token creates a Json Web Token using the jose2go library.
// It takes a payload as a json string and the path to a public certificate.
// It returns the token as a string and any errors that may have occurred.
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

// VerifyToken verifies a Json Web Token using a given private key.
// It takes the token and a path to the private key file as strings.
// It returns the payload as a json string, the headers as a map[string]interface{},
// and any errors that may have occurred.
func VerifyToken(token, privKey string) (string, map[string]interface{}, error) {

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
