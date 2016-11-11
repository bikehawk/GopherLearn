package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 1000; i++ {
		if i%10 == 0 {
			fmt.Println(i * 10)
		}
	}
}
