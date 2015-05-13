package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 15; i++ {
		for j := 1; j <= 15; j++ {
			fmt.Printf("%d\t", i*j)
		}
		fmt.Println()
	}
}
