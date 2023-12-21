package main

import "fmt"

func main() {
	var a int = 1
	var b float64 = 2.1
	//Invalid operation: a + b (mismatched types int and float64)
	// fmt.Println(a + b)

	fmt.Println(float64(a) + b)

}
