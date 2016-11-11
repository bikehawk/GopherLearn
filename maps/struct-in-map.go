package main

import (
	"fmt"
)

type T struct {
	T1 string
	T2 string
}

var m = map[string]*T{
	"test1": &T{
		T1: "test string 1",
		T2: "test string 2",
	},
	"test2": &T{
		T1: "test string 3",
		T2: "test string 4",
	},
}

func main() {
	fmt.Println(m["test1"].T2)
}
