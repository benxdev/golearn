package main

import (
	"fmt"
	)
 
var c = true

func main() {
	var a int32 = 56
	var b int64 = 56
	var c = a + b
	a = b
	fmt.Println(c)
}
