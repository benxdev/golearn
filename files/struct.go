package main

import (
	"fmt"
)

func main() {
	type HumanBody struct {
		Head, Arm            string
		Height, Weight, Span float64
		SkinColour           string
	}
	me := HumanBody{
		"big",
		"strong",
		172.5,
		83.1,
		55.5,
		"light",
	}

	fmt.Println(me)
	newHeight := &me.Height // pointer to the height field on me struct with type HumanBody
	fmt.Printf("%T %.2f\n", newHeight, *newHeight)
	fmt.Printf("%T\n", me)

	john := &HumanBody{
		"bigger",
		"weak af",
		169.8,
		75.7,
		50.2,
		"midnight black",
	}

	fmt.Println(john)
	johnNewHeight := &john.Height // pointer to the height field on john struct with type *HumanBody
	fmt.Printf("%T %.2f\n", johnNewHeight, *johnNewHeight)
	fmt.Printf("%T\n", john)
}
