package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("MYVAR", os.Getenv("MYVAR"))
	if os.Getenv("MYVAR") == "" {
		fmt.Println("MYVAR is not set; setting MYVAR to test.")
		os.Setenv("MYVAR", "test")
	}
	fmt.Println("MYVAR:", os.Getenv("MYVAR"))
}
