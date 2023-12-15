package main

import "fmt"

func sayGreeting(n string) {
	fmt.Printf("Greetings, %v\n", n)
}

func sayBye(n string) {
	fmt.Printf("Bye, %v\n", n)
}

func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func main() {

	sayGreeting("Mario")
	sayGreeting("Luigi")

	sayBye("Mario")

	cycleNames([]string{"Cloud", "Sun", "Moon"}, sayGreeting)

}
