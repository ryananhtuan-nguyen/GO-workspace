//ARRAY AND SLICES

package main

import "fmt"

func main() {
	//declaring an array
	//arrays in GO has FIXED length
	// var ages [3]int = [3]int{2, 3, 4}

	var ages = [3]int{20, 30, 40}

	names := [4]string{"Ryan", "Ryan1", "Ryan2", "Ryan3"}

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

}
