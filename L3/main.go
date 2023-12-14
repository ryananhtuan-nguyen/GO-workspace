//ARRAY AND SLICES

package main

import "fmt"

func main() {
	//declaring an array
	//arrays in GO has FIXED length
	// var ages [3]int = [3]int{2, 3, 4}

	var ages = [3]int{20, 30, 40}

	fmt.Println(ages, len(ages))

}
