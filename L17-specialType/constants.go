package main

import "fmt"

func main() {

	//Constants are declared like variables, but with keyword const
	//can be character, string, bool, or numeric
	//can not be declared using :=
	const PI = 3.14159

	//Cannot assigned to constant:
	// PI := 123

	fmt.Println(PI)

	//uintptr - Unsigned Int Pointer

	/*
		key-value
		key = uintptr -> value
		map<key,object>
		value = object(*point)

		key=(*point) -> utintptr
		map.put(key, value)
		List of address, each address points to a value.
	*/
}
