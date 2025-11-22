package main

import (
	"fmt"
	"strings"
)

var st1 = []string{
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
        "Akasha",    
        "Planned",      
        "Crystal",   
        "Dazzle",    
        "Enigma",    
        "Faceless",  
        "Gondar",    
        "Rattlesnake",    
        "Invoker",   
        "Juggernaut",
        "Kunkka",  
		"ccc",  
    }

type I interface {
	showShortNames() []string
	rmvDoubleLetterWords() []string
	collateDoubleLetterWords() []string
}

type myList struct {
	Names []string
}

func (ml myList) showShortNames() []string {
	output := []string{}
	for _, name := range ml.Names {
		r := []rune(name) // convert each looped string to a slice of runes
		if len(r) <= 6 { // checks each word length by number of runes, not number of bytes
			output = append(output, name)
		} 
	}
	return output
}

func hasDoubleLetter(word string) bool {
	r := []rune(strings.ToLower(word)) // convert current looped string to a slice of runes, also ensures case insensitivity.
	for i := 0; i < len(r)-1; i++ { // loops through the entire slice of runes
		if r[i] == r[i+1] { // checks for repeated runes in the slice of runes
			return true
		}
	}
	return false
}

func (ml myList) collateDoubleLetterWords() []string {
	collated := []string {}
	for _, word := range ml.Names { // loops through the entire slice of strings
		if hasDoubleLetter(word) {
			collated = append(collated, word)
		}
	}
	return collated
}
func (ml myList) rmvDoubleLetterWords() []string {
	collated := []string {}
	for _, word := range ml.Names { // loops through the entire slice of strings
		if !hasDoubleLetter(word) {
			collated = append(collated, word)
		}
	}
	return collated
}

func main() {
	list := myList{Names: st2,}
	// list2 := myList {Names: st2,}
	var i I = list
	fmt.Printf("The double letter words are: %v\n", i.collateDoubleLetterWords())
	fmt.Printf("The non double letter words are: %v\n", i.rmvDoubleLetterWords())
	fmt.Printf("The short names here: %v\n", i.showShortNames())
}
