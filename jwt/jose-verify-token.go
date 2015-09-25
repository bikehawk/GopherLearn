package main

import (
	"fmt"
	"github.com/dvsekhvalnov/jose2go"
	"github.com/dvsekhvalnov/jose2go/keys/rsa"
	"io/ioutil"
)

func main() {

	token := "eyJhbGciOiJSU0EtT0FFUCIsImVuYyI6IkEyNTZHQ00ifQ.dfkEcCuHII5a-lpvGRcnxHxbqvqDzX2yRcUiJh0p04tdCLPmN2FvaKKOWAAaG8-ag5VAt0uUsTRoRsJPx2dI5MvrkP1_yn4bz0b5JYqbCjOZwG_UEtEhUOfBRIQheuT_ydy6QolwSxrS2J0l3EmJS98XZ8yjGdxcaKsylIS7BLfBN8a51YDlKxNBQkHwa6oRbT7du3wHvzbXjVvw2__hJSdnHm9T1lf8ayCLgJne59GEuSdaoEAc-Sv_4hBEj1i4L98q-ewM2HnhH3xmFLqUWjlzzK4vSpR2_k2giwwtBxfPPWERMl8wOmCgvlxl6q9Njh_Jgh5BgQktSFI40EvQYA.c5sLigtCIQlnbQYJ.fKRVMkD2dOH4_GLZLv1iRjSx.0Z45LEQOSbEhIE6pd2gqvw"

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
	}
}
