package main

import (
	"fmt"
)

var PI float64 = 3.14

type Race struct {
	Race, Nationality string
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

var i int = 4
var j int = 65

func f00(i int, st []string) (int, string) {
	return i, st[5]
}

func f01(j int, PI float64) (int, int) {
	shortPI := int(PI)
	return j, shortPI
}

func f(func(int, []string) (int, string), func(int, float64) (int, int)) (int, string, int) {
	_, shortPI := f01(j, PI) // because shortPI here is outside the scope of f.
	return i + j, st[5], shortPI
}

func main() {
	fmt.Println(f(f00, f01))
}
