//GO is a pass-by-value language

package main

import "fmt"

func updateName(x string) string {
	x = "shoestrings fries"
	return x
}

func updateMenu(y map[string]float64) {
	y["coffee"] = 2.99
}

func main() {

	//Non-pointer values, straight forward values
	//FIRST GROUP: types -> strings, ints, bools, floats, arrays, structs
	name := "tifa"

	//This will NOT change value of name
	// updateName(name)
	name = updateName(name)

	fmt.Println(name)

	//Pointer wrapper values: value that store a pointer referencing to the actual value store in another place
	//SECOND GROUP types -> slices, maps, functions

	menu := map[string]float64{
		"pie":       5.95,
		"ice cream": 3.99,
	}

	//This WILL change the original
	updateMenu(menu)

	fmt.Println(menu)
}
