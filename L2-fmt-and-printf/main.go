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

	//Printf ( formatted strings )  %_ = format specifier
	//%v for variable
	fmt.Printf("My age is %v and my name is %v \n", age, name)
	//%q string
	fmt.Printf("My age is %q and my name is %q \n", age, name)
	//%T type
	fmt.Printf("Age is of type %T \n", age)
	//%f FLOAT
	fmt.Printf("You scored %0.1f points! \n", 225.55)

	//Sprintf (save formatted strings)
	var str = fmt.Sprintf("My age is %v and my name is %v \n", age, name)
	fmt.Println("The saved string is:", str)

}
