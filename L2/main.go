package main

import "fmt"

func main() {

	age := 15
	name := "Ryan"

	// Print, same line
	// fmt.Print("Hello, ")
	// fmt.Print("World! \n")
	// fmt.Print("New Line! \n")

	// fmt.Println("Another new line")
	// fmt.Println("Bye !")

	fmt.Println("My age is", age, "and my name is", name)

	//Printf ( formatted strings )
	fmt.Printf("My age is %v and my name is %v", age, name)

}
