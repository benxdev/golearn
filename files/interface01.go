package main

import (
	"fmt"
	"strings"
)

type sType []string

type Slicer interface {
	enstring() string
	search(string) string	// because search method takes a string arg.
}

// merges all string elements in any []string.
func (s sType) enstring() string {
	merge := strings.Join(s, " ")
	return merge
}
func (s sType) search(key string) string {
	for _, v := range s {
		if v == key {
			return key + " exists."
		}
	}
	return key + " not found."
}

func main() {
	var ser Slicer = sType(Sst1) // Sst1 defined in data.go
	fmt.Println(ser.search("Juggernaut")) 
	fmt.Println(ser.search("Abiliwa")) 
}
