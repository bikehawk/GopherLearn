package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("MYENVAR:", os.Getenv("MYENVAR"))
	fmt.Println("Setting MYENVAR...")
	os.Setenv("MYENVAR", "set")
	fmt.Println("MYENVAR:", os.Getenv("MYENVAR"))
}
