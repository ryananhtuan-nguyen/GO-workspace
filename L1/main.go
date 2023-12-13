package main

import "fmt"

// func main() {
// 	fmt.Println("Hello world!")
// }

func main() {

	//strings
	var nameOne string = "mario"

	var nameTwo = "Ryan"

	var nameThree string

	// nameThree = nameOne
	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "orange"
	nameThree = "browser"

	fmt.Println(nameOne, nameTwo, nameThree)

	nameFour := "Fourth name" //This can NOT be used outside of a function

	fmt.Println(nameFour)

	//ints

	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)
}
