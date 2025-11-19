package main
import (
	"fmt"
)

var PI float64 = 3.14
type Race struct {
	Race, Nationality string
}
var st = []string{
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

func createFibSequence(first, second int) func() int {
	a, b := first, second
	return func() int {
		result := a
		a, b = b, a+b
		return result
	}
}

func main() {
	fib := createFibSequence(17, 23)
	for i := 0; i < 20; i++ {
		fmt.Println(fib())
	}

}