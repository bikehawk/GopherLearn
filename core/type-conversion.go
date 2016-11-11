package main

import "fmt"

func main() {
	uint16val := 65000
	fmt.Printf("uint16val: %d\n", uint16val)

	int16val := int16(uint16val)
	fmt.Printf("int16val: %d\n", int16val)
}
