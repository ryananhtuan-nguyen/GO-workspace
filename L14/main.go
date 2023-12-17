package main

import "fmt"

//STRUCTS & CUSTOM TYPES

func main() {
	mybill := newBill("Ryan's bill")

	mybill.updateTip(10)

	fmt.Println(mybill.format())
}
