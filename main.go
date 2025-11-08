package main

import (
	"fmt"
)

func main() {
	oddsArr := [10]int{3, 9, 69, 55, 4565, 656, 68, 898, 658, 12}
	fmt.Println(oddsArr)

	oddsSlc := oddsArr[1:(len(oddsArr)-2)]

	fmt.Printf("Length: %d, Value: %v, Type: %T", len(oddsSlc), oddsSlc, oddsSlc)
}