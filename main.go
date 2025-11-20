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
	func (r Race) mthd()string {
		return "I speak fucking Chinese!"
	} 
	
func main() {
	nr := Race{"Asian", "Chinese"}
	fmt.Println(nr.mthd())
}