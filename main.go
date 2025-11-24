package main

import (
	"fmt"
	"strings"
)

type Matrix [][]string
type TTT struct {
	PlayerOne Player
	PlayerTwo Player
	History []Matrix
	Moves int
}
type Player struct {
	Player string
}


func factoryMatrix() [][]string {
	// display the matrix
	fmt.Print("original matrix:\n")
	matrix := Matrix{{"-", "-", "-"}, {"-", "-", "-"}, {"-", "-", "-"}}
	for _, member := range matrix {
		fmt.Printf("%s\n", strings.Join(member, " "))
	}

	var userPosition string
	fmt.Print("\nset position to play:\n")
	fmt.Scan(&userPosition)

	// values must be either X or O (not strictly set yet)
	fmt.Print("set play value:\n")
	var playerValue string
	fmt.Scan(&playerValue)
	playerValue = strings.ToUpper(playerValue)

	// sets the return integers of checkPositionInput to row and col matrix index positions
	row, col := checkPositionInput(userPosition) // e.g:  checkPositionInput("TL")[row col]
	matrix[row][col] = playerValue

	fmt.Print("\ncurrent matrix:\n")
	for _, member := range matrix {
		fmt.Printf("%s\n", strings.Join(member, " "))
	}

	return matrix
}

func main() {
	for i := 0; i < 5; i++ {

	}
	factoryMatrix()

}

func checkPositionInput(s string) (int, int) {
	switch s {
	case "tl", "TL", "top left", "TOP LEFT":
		return 0, 0
	case "tm", "TM":
		return 0, 1
	case "tp", "TP":
		return 0, 2
	case "cl", "CL":
		return 1, 0
	case "cc", "CC":
		return 1, 1
	case "cr", "CR":
		return 1, 2
	case "bl", "BL":
		return 2, 0
	case "bc", "BC":
		return 2, 1
	case "br", "BR":
		return 2, 2
	default:
	}
	return 0, 0
}
