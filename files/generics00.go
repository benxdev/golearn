package main

import "fmt"

func Index[T comparable](s []T, x T) int { // T must support the == and != operators.
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

func main() {
	si := []int{23, 324, 434, 5, 22, 0, 9, 567}
	fmt.Println(Index(si, 5))

	ss := []string{"fog", "dog", "ekweremadu", "cynical", "abuja"}
	fmt.Println(Index(ss, "abuja"))
}
