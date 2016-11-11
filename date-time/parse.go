package main

import (
	"fmt"
	"time"
)

func main() {
	const dateFormat = "Jan 2, 2006 (UTC)"
	t, err := time.Parse(dateFormat, "Mar 18, 1976 (UTC)")

	if err != nil {
		fmt.Printf("error: %s\n", err)
	}
	fmt.Println(t)
}
