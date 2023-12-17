package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//get input

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

//create Bill

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create a new bill name: ", reader)

	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)

	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose option (a - add item, s - save bill, t - add tip): ", reader)

	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)

		fmt.Println(name, price)
	case "t":
		tip, _ := getInput("Enter tip amount ($): ", reader)

		fmt.Println(tip)
	case "s":
		fmt.Println("You chose s")
	default:
		fmt.Println("That was not a valid option...")
		promptOptions(b)
	}

}

// MAIN FUNC
func main() {
	mybill := createBill()
	promptOptions(mybill)
}
