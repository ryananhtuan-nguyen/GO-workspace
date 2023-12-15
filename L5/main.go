package main

import "fmt"

func main() {
	//FOR LOOP

	// x := 0
	// for x < 5 {
	// 	fmt.Println("value of x is:", x)
	// 	x++
	// }

	//TRADITIONAL FOR LOOP
	// for i := 0; i < 5; i++ {
	// 	fmt.Println("value of i is:", i)

	// }

	//LOOPING THROUGH SLICES

	names := []string{"mario", "moria", "yoda", "watermelon"}

	//traditional For
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])

	}

	//for...in like js
	//switch variable to _ if not needed
	//for _, value
	for index, value := range names {
		fmt.Printf("Value at index %v is %v \n", index, value)

		//THIS WONT UPDATE REAL VALUE
		value = "something"
	}

	fmt.Println(names)

}
