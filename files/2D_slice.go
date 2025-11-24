package main

import (
	"fmt"
	"strings"
)

func main() {	
	board := [][]string {
		{"-", "-", "-"},
		{"-", "-", "-"},
		{"-", "-", "-"},
	}
	for _, members := range board{
		fmt.Printf("%s\n", strings.Join(members, " "))
	}

	board[1][1] = "XX"
	board[2][0] = "00"
	board[0][2] = "CC"

	for _, members := range board{
		fmt.Printf("%s\n", strings.Join(members, " "))
	}
}