package main

import (
	"fmt"
)
type Ser interface {
	fullNamer()
}

type Person struct {
	FirstName string
	LastName string
}

func (p Person) fullNamer() {
	fmt.Print("first name = ")
	fmt.Scanln(&p.FirstName)
	fmt.Print("lastname = ")
	fmt.Scanln(&p.LastName)
	fmt.Println("fullname = " + p.FirstName + " " + p.LastName) 
}

func main() {
	var p Person
	var ser Ser = p // OR ser := Ser(p)
	ser.fullNamer()
}
