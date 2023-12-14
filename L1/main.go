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

	//bits and memory
	// var numOne int8 = 25
	// var numTwo int8 = -128
	// var numThree uint8 = 225

	//FLOAT
	var scoreOne float32 = -22.5
	var scoreTwo float64 = 32487234982374.092382487
	scoreThree := 1.5

}
