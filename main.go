package main

import (
	"fmt"
)

var (
	a = 930
	b = 250.0
)

func main() {
	// a := 4
	fmt.Printf("Type of a is %T, value of a is %[1]v\n", a)
	fmt.Printf("Type of b is %T, value of b is %[1]v\n", b)

	c:= a + int(b)
	fmt.Println(c)
}