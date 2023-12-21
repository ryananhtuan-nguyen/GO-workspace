package main

import "fmt"

func main() {
	//RUNE
	var myString string = "biết"

	runes := []rune(myString)

	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c", runes[i])
	}

	fmt.Println("============")
	var myRune rune = 'A'
	fmt.Printf("%T", myRune)

	fmt.Println("============")
	var mRune rune = '🔴'
	fmt.Printf("%c", mRune)

}
