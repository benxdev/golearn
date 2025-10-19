package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5}
	y := make([]int, 2)
	copy(y, x[3:])
	fmt.Println("y = ", y)
}
