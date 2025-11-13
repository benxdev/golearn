package main

import (
	"fmt"
)

func main() {	
	var nums []int
	for i := 0; i < 5000; i++ {
		nums = append(nums, i)
		fmt.Printf("Length=%d, Capacity=%d\n", len(nums), cap(nums))
	}
}



// s := []string{
	// 	"John",
	// 	"Okafor",
	// 	"Ibekwe",
	// 	"Merlin",
	// 	"Paul",
	// 	"Juggernaut",
	// 	"Heathrow",
	// 	"Presbitarian",
	// }