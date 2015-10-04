package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("MYENVAR: %#v\n", os.Getenv("MYENVAR"))
	fmt.Println("Setting MYENVAR...")
	os.Setenv("MYENVAR", "set")
	fmt.Printf("MYENVAR: %#v\n", os.Getenv("MYENVAR"))
}
