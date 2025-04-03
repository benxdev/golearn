package main 
import "fmt"

const MAX = 100
func getPoint(x int, y int) (int,int) {
	return x, y
}
func main() {

	if 69 < 0 {
		x, _ := getPoint(3, 5)
		fmt.Println(x)
	} else {
		fmt.Println("not required")
	}
	
}