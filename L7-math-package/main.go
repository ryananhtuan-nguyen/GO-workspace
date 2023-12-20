package main

import (
	"fmt"
	"math"
)

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

func circleAre(r float64) float64 {
	return math.Pi * r * r
}

func main() {

	// sayGreeting("Mario")
	// sayGreeting("Luigi")

	// sayBye("Mario")

	cycleNames([]string{"Cloud", "Sun", "Moon"}, sayGreeting)
	cycleNames([]string{"Cloud", "Sun", "Moon"}, sayBye)

	a1 := circleAre(10.5)
	a2 := circleAre(15)

	fmt.Println(a1, a2)
	fmt.Printf("Circle 1 is %0.3f and circle2 is %0.3f \n", a1, a2)

}
