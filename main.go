// package main

// import (
// 	"fmt"
// )

// var x int = 344

// func main() {
// 	var x = make([]int, 5, 50)
// 	x = append(x, 67)
// 	x[1] = 69
// 	fmt.Println(x, len(x), cap(x))

// }


package main
import "fmt"

func main() {
	x := []string{"a", "b", "c", "d"}
	y := x[:2]	//[a, b]
	z := x[1:]	//[b, c, d]
	x[1] = "y"	//x = [a, y, c, d]
	y[0] = "x"	//y = []
	z[1] = "z"
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
}
