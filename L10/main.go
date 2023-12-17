package main

import "fmt"

func main() {

	menu := map[string]float64{
		"soup":    4.99,
		"pie":     7.99,
		"salad":   6.99,
		"pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["pudding"])

	//LOOPING THROGH maps
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	//ints as key type
	phonebook := map[int]string{
		0123:  "Moria",
		4567:  "Abc",
		89910: "cba",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[0123])

	phonebook[4567] = "Different Name"

	fmt.Println(phonebook)

	phonebook[89910] = "Something different"

	fmt.Println(phonebook)
}
