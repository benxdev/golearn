//EXAMPLE 1
package main
import (
	"fmt"
	)

func main() {
	age := 57
	agePtr := &age
	fmt.Println("Age is: ", *agePtr)

	*agePtr = 690
	fmt.Println("Age is now: ", *agePtr)

	fmt.Print("New value now:\n")
	f := 567
	*agePtr = f
	fmt.Println("Age has now become: ", *agePtr)
}


// EXAMPLE 2:
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var a int
// 	fmt.Print("Enter the value of a\n")
// 	fmt.Scanf("%d", &a)
// 	var p *int // pointer p (variable that stores the memory address of another variable)
// 	p = &a     // pointer p stores the "memory address" of variable a
// 	var c = *p // c is declared and takes on the value of variable (a) that the pointer p is pointing to in memory
// 	fmt.Printf("c copied as a: %d", c)
// }

