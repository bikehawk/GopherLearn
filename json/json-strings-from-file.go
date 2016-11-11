package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Values []string

func main() {
	c, err := ioutil.ReadFile("./example-json-file.json")
	if err != nil {
		log.Fatal("Could not open json file.")
	}

	dec := json.NewDecoder(bytes.NewReader(c))

	var v Values

	dec.Decode(&v)

	fmt.Printf("%#v\n", v)
}
