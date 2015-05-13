package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "searchtermone|search term two|search_term-three"
	st := strings.Split(s, "|")
	fmt.Printf("%#v\n", st)
}
