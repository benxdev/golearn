package main
import (
	"fmt"
)

func factory() func(x int) int {
	n := 23
	return func (x int)int {
		n = n + x
			return n
	}
}

func main() {
	newFac := factory()
	anoFac := factory()

	fmt.Println(newFac(7))
	fmt.Println(anoFac(4-7))
	fmt.Println(newFac(7))
	fmt.Println(anoFac(4*7))
	fmt.Println(newFac(7))
	fmt.Println(anoFac(4*7))
}


// Fibonacci sequence implemented using closures:

// func createFibSequence(first, second int) func() int {
// 	a, b := first, second
// 	return func() int {
// 		result := a
// 		a, b = b, a+b
// 		return result
// 	}
// }

// func main() {
// 	fib := createFibSequence(17, 23)
// 	for i := 0; i < 20; i++ {
// 		fmt.Println(fib())
// 	}

// }