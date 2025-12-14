package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	)

func do(i interface{}) {
	switch v := i.(type) {
		case int: 
			fmt.Print(v, " is an int")
		case string:
			fmt.Print(v, " is a string")
		case bool:
			fmt.Print(v, " is a boolean")
		default:
			fmt.Print(v, " is currently undetermind")
	}
}

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter something: ")
	
	if scanner.Scan() {
		s := scanner.Text()
		
		if num, err := strconv.Atoi(s); err == nil {
			do(num)
		} else if s == "true" || s == "false" {
			if b, err := strconv.ParseBool(s); err == nil {
				do(b)
			}
		} else {
			do(s)
		}
	}
}
