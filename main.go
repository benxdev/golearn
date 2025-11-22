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

var st2 = []string{
        "Akasha",    // 6
        "Bane",      // 5
        "Crystal",   // 7
        "Dazzle",    // 7
        "Enigma",    // 6
        "Faceless",  // 8
        "Gondar",    // 6
        "Huskar",    // 6
        "Invoker",   // 7
        "Juggernaut",// 9
        "Kunkka",    // 6
    }

type I interface {
	showName() string
}

type myList struct {
	Names []string
}


func (sh myList) showName() string {
	for i:=0; i<len(sh.Names); i++ {
		if len(sh.Names[i]) <= 6 {
			fmt.Println(sh.Names[i])
		} 
	}
	return "Too Large a string"
}

func (sh myList) removeSomeEntries() []string {
	
}

func main() {
	list := myList {
		Names: st,
	}
	list2 := myList {
		Names: st2,
	}


	var i, j I = list, list2
	i.showName()
	j.showName()

}
