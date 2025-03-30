package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Print("Enter the value of a\n")
	fmt.Scanf("%d", &a)
	var p *int 	// pointer p (variable that stores the memory address of another variable)
	p = &a 		// pointer p stores the "memory address" of variable a
	var c = *p 	// c is declared and takes on the value of variable (a) that the pointer p is pointing to in memory
	fmt.Printf("c copied as a: %d", c)
}