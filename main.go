package main

import (
	"fmt"
)

func pString(c string) {
	fmt.Printf("this input is a %T type", c)
}
func pInt(c int) {
	fmt.Printf("this input is a %T type", c)
}

func main() {
	var i interface{}
	fmt.Scan(&i)

	switch v := i.(type) {
	case string:
		pString(v)
	case int:
		pInt(v)
	default:
		fmt.Print("invalid input type!")
	// 	s, ok := i.(int16)
	// 	fmt.Println(s, ok)
	}
}
