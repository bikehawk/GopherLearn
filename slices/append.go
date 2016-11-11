package main

import (
	"fmt"
)

func main() {
	slc := []string{"item1", "item2"}
	str := "item0"

	slc = append(slc, 0)

	fmt.Printf("%#v\n", slc)
}
