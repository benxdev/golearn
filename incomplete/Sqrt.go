package main

import (
	"fmt"
)

// Implementing the math.Sqrt() function from the Standard Library
func Sqrt(x float64) (root float64) {

	if z := 1.0; z*z == x {
		root = 1.0
		return root
	} else {
		for z = 1.0; z < x; z++ {
			root = ((z * z) - x) / 2 * z

		}
	}
	return root
}

func main() {
	var x float64
	fmt.Print("Enter number to find square root:")
	fmt.Scanf("%f", &x)
	root := Sqrt(x)
	fmt.Printf("The square root of %v is %v", x, root)
}
