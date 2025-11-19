package main
import (
	"fmt"
)

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
