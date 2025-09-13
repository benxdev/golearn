package main

import (
	"fmt"
)

var x int = 344

func main() {
	var x = make([]int, 5, 50)
	x = append(x, 67)
	x[10] = 69
	fmt.Println(x, len(x), cap(x))

}
