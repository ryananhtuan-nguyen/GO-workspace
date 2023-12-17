package main

import "fmt"

func updateName(x *string) {
	*x = "fries"
}

func main() {
	name := "tifa"

	// updateName(name)

	// fmt.Println("memory address of name is:", &name)

	m := &name

	// m and name are pointing to the same memory address
	fmt.Println("Memory address of m is:", m)
	fmt.Println("Value at memory address:", *m)

	//this WILL actually change the value
	updateName(m)

	fmt.Println(name)
}

/*
|--name---|----m----|
|  0x001  |  0x002  |
|---------|---------|
| "tifa"  | p0x001  |
|---------|---------|
*/
