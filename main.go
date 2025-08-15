package main

import (
	"fmt"
	"math"
)

const TestDigit = 69

func main(){

	var b byte 			= 255
	var smallI int32 	= 2147483647
	var bigI uint64 	= 18446744073709551615

	b += 1
	smallI += 1
	bigI += 1

	fmt.Println(b, smallI, bigI)
	var bigNum uint64 = math.MaxUint64
	fmt.Println(bigNum)

}