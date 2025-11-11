package main

import "fmt"

func main() {
	x := make([]string, 0, 5)			// make() is used to initialize Reference types only (Channels, Maps, Slices, ...)
	x = append(x, "a", "b", "c", "d")
	y := x[:2]                          // y = a,b
	z := x[2:]                          // z = c,d
	fmt.Println(cap(x), cap(y), cap(z)) // 5 5 3 | A slice’s capacity = total backing capacity - starting index of slice
	y = append(y, "i", "j", "k")        // y = a,b,i,j,k
	x = append(x, "x")                  // x = a,b,i,j,x
	z = append(z, "y")                  // z = i,j,y | y = a,b,i,j,y (y also now changes as a result of shared memory between the sliced slices)
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)

}
