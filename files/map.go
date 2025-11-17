package main

import (
	"fmt"
)
var PI float64 = 3.14

type HumanBody struct {
	HeadShape, ArmSize       string
	Height, Weight, WingSpan float64
	SkinColourType           string
}

var st = []string{
	"John",
	"Okafor",
	"Ibekwe",
	"Merlin",
	"Paul",
	"Juggernaut",
	"Heathrow",
	"Presbitarian",
	"Onome",
	"Odogwu Bitters",
	"Liquidity",
}

func main() {
	m0 := make(map[string]*HumanBody)  // struct fields here are modifiable because the map value is a pointer.
	m0["Ejike"] = &HumanBody{
		"Oval", "Big", 172.5, 83.2, 185.9, "Light skin",
	}
	myHeight := m0["Ejike"].Height
	hd := st[5]
	m0["Ejike"].HeadShape = hd // using pointer allows this.
	fmt.Printf("height = %.2f, head size = %s\n\n", myHeight, hd)

	m01 := make(map[string]HumanBody) // struct fields here are non-modifiable.
	m01["Bibi"] = HumanBody{
		"Round", "Small", 170.0, 79.3, 174.5, "Light skin",
	}

	// Regular Maps //
	m1 := make(map[string]int)
	m2 := make(map[string]string)
	m1["my age"] = 26
	m2["boy"] = "A biological male human child"
	fmt.Printf("Ejike is %s aged %d years old\n", m2["boy"], m1["my age"])

	m3 := map[string]string {"girl": "A biological female human child",}
	fmt.Printf("%v\n", m3["girl"])

	m4 := map[string]float64 {
		"PI": 3.14,
	}
	fmt.Printf("value of PI = %v\n", m4["PI"])

}
