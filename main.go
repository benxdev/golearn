package main
import (
	"fmt"
)
//anonymus structs in go

type man struct {
	name string
	age int
	race string
 
}

type oneman man struct {
	name = "Ejike"
	age = 45
	race = "nigger"
}

func main() {
	var newUser user

	newUser.name = "EJIKE\n"
	newUser.number = 9809800

	fmt.Print(newUser.name, newUser.number)
}