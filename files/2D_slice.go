package main

import (
	"fmt"
	"strings"
)

func main() {	
	board := [][]string {
		[]string {"-", "-", "-"},
		[]string {"-", "-", "-"},
		[]string {"-", "-", "-"},
	}
	for i:=0; i<len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	board[1][1] = "XX"
	board[2][0] = "00"
	board[0][2] = "CC"

	for i:=0; i<len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}