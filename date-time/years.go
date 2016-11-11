package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	then := now.AddDate(-14, 0, 0)

	y, _, _ := then.Date()

	var a [76]int

	for i, j := (y - 75), 0; i <= y; i, j = i+1, j+1 {
		a[j] = i
	}

	fmt.Println(a)
}
