package main

import (
	"fmt"
)

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

type I interface {
	doIt() []string
}

type myList struct {
	Names []string
}


func (m *myList.Names) doIt(c []string) {
	m = &myList.Names
	for i:=0; i<5; i++ {
		fmt.Println(m)
	}
}

func main() {
	 list :=  myList {st}
	var i I = 
	c := &myList{st}

}
