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

func factory() func(x int) int {
	n := 23

	return func (x int)int {
		n = n + x
		return n
	}
}

func main() {
	newFac := factory()
	anoFac := factory()

	fmt.Println(newFac(7))
	fmt.Println(anoFac(4-7))
	fmt.Println(newFac(7))
	fmt.Println(anoFac(4*7))
	fmt.Println(newFac(7))
	fmt.Println(anoFac(4*7))
}
