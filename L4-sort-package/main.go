package main

import (
	"fmt"
	// "strings"
	"sort"
)

func main() {
	// greeting := "Hello there world!"

	// fmt.Println(strings.Contains(greeting, "Hello"))
	// fmt.Println(strings.ReplaceAll(greeting, "Hello", "Hi"))
	// fmt.Println("Original string value =", greeting)
	// fmt.Println(strings.ToUpper(greeting))

	// fmt.Println(strings.Index(greeting, "ll"))

	//Split return an array
	// fmt.Println(strings.Split(greeting, ""))

	ages := []int{45, 20, 35, 30, 75, 60, 70, 25}

	//sort mutate the original variable
	// fmt.Println(ages)

	//SORT NUMBER
	sort.Ints(ages)
	fmt.Println(ages)

	//Sort.searchInt: return the index of item in the slice
	//if length is not specified it will return len
	index := sort.SearchInts(ages, 30)
	fmt.Println(index)

	names := []string{"aryan0", "dryan1", "cryan2", "bryan3"}

	//SORT STRING
	// sort.Strings(names)
	fmt.Println(names)

	//SEARCH STRING
	fmt.Println(sort.SearchStrings(names, "cryan2"))
}
