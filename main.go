package main

import "fmt"

type car struct {
	Make string
	Model string
	Height int
	Width int
	Frontwheel Wheel
	Backwheel Wheel
	Category KindOfVehicle
}
type KindOfVehicle struct {
	Kind string
	Power float64	
}
type Wheel struct {
	Radius int
	Material string
}

func main () {
	myCar := car{}
	var theType := myCar.Category.Kind
}