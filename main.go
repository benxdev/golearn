package main

import (
	"fmt"
)
var PI float64 = 3.14

type Race struct {
	Race, Nationality  string
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

var f00 = func(i int, st []string)(int, string) {
	return i, st[5]
}

var f01 = func(j int, PI float64)(int, int) {
	shortPI := int(PI)
	return j, shortPI
}

var f = func(func(int, []string)(int, string), func(int, float64)(int, int)) (int, string, int) {
	_, shortPI := f01(j, PI)
	return i + j, st[5], shortPI
}

func main() {
	fmt.Println(f(f00, f01))
}

