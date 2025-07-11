package main

import (
	"fmt"
)

func mathFunc(a, b int) (prod int, sum int) {
	prod = a * b
	sum = a + b
	return
}
func main() {
	var num1, num2 int
	fmt.Print("Enter Num 1: ")
	fmt.Scanf("%d\n", &num1)
	fmt.Print("Enter Num 2: ")
	fmt.Scanf("%d\n", &num2)
	var prod, sum int = mathFunc(num1, num2)
	fmt.Printf("prod: %d sum: %d\n", prod, sum)
}
