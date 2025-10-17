package main

import "fmt"

func main() {
	x := []string{"a", "b", "c", "d"}
	y := x[:2] // y = a,b
	z := x[1:] // z = b,c,d
	x[1] = "y" // x = a,y,c,d | y = a,y | z = y,c,d
	y[0] = "x" // x = x,y,c,d | y = x,y | z = y,c,d
	z[1] = "z" // x = x,y,z,d | y = x,y | z = y,z,d
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)

}
