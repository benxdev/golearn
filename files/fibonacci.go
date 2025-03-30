package main
import (
	"fmt"
	// "math"
)
func fibonacci(n int) int {
	if n < 0 {
		return 0
	} else if n == 1 {
			return 1 
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

func main() {
	var n int
	fmt.Print("enter a number to calculate sequence: ")
	fmt.Scanf("%d", &n)
	fmt.Printf("the sequence at position %d is %d\n", n, fibonacci(n))
}