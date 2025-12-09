package main 

import (
	"fmt"
	"strings"
)
type sType []string

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
	fmt.Println(sType(Sst1).enstring())     // Sst1 defined in data.go (may be outside of scope)
	fmt.Println(sType(Sst1).search("okpe")) // Sst1 defined in data.go (may be outside of scope)
}
