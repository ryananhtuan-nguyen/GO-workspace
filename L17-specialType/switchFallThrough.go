package main

import "fmt"

func main() {
	//Fallthrough
	//continue to check for other cases after found a match

	number := 10
	switch number {
	case 1:
		fmt.Println("Number = 1")
		fallthrough
	case 10:
		fmt.Println("Number = 10")
		fallthrough
	case 2:
		fmt.Println("Number = 2")
		fallthrough
	case 3:
		fmt.Println("Number = 3")
	}
}
