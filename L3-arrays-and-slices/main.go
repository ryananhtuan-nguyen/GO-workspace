//ARRAY AND SLICES

package main

import "fmt"

func main() {
	//declaring an array
	//arrays in GO has FIXED length
	// var ages [3]int = [3]int{2, 3, 4}

	var ages = [3]int{20, 30, 40}

	names := [4]string{"Ryan", "Ryan1", "Ryan2", "Ryan3"}
	names[1] = "RyanN"

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	//slices (use arrays under the hood)

	var scores = []int{100, 50, 60}
	scores[2] = 25
	fmt.Println(scores)
	scores = append(scores, 200)
	fmt.Println(scores, len(scores))

	//slice ranges, endIndex exclusive
	rangeOne := names[1:3]
	//omit end to go all the way to the end
	rangeTwo := names[2:]
	//omit start vice versa, end Index exclusive
	rangeThree := names[:3]
	fmt.Println(rangeOne)
	fmt.Println(rangeTwo)
	fmt.Println(rangeThree)

	rangeOne = append(rangeOne, "Kappa")
	fmt.Println(rangeOne)
}
