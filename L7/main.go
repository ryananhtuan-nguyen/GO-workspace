package main

import "fmt"

func sayGreeting(n string) {
	fmt.Printf("Greetings, %v\n", n)
}

func sayBye(n string) {
	fmt.Printf("Bye, %v\n", n)
}
func main() {

	sayGreeting("Mario")
	sayGreeting("Luigi")

	sayBye("Mario")

}
