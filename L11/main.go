//GO is a pass-by-value language

package main

import "fmt"

func updateName(x string) {
	x = "shoestrings fries"
}

func main() {
	name := "tifa"

	updateName(name)

	fmt.Println(name)
}
