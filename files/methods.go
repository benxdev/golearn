package main

import (
	"fmt"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) { // uses pointer receiver (v *Vertex). original struct mutates after method computes.
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Vertex) Scales(f float64) { // uses value receiver (v Vertex). original struct does not mutate after method computes.
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Scaless(f float64) { // same as Scale()
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	nv := Vertex{3, 4} 
	nv.Scale(2)
	fmt.Println(nv) // {6, 8},  mutation occurs as a result of pointer receiver being used in Scale()

	nv2 := Vertex{3, 4} // {3, 4},  mutation does not occur as a result of value receiver being used in Scales()
	nv2.Scales(2)
	fmt.Println(nv2)

	nv3 := Vertex{3, 4}
	nv3.Scaless(2)
	fmt.Println(nv3)
}
