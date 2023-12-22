// Given an array of integers your solution should find the smallest integer.

// For example:

// Given [34, 15, 88, 2] your solution will return 2
// Given [34, -345, -1, 100] your solution will return -345
// You can assume, for the purpose of this kata, that the supplied array will not be empty.

package kata

import "sort"

// THIS FUNCTION O(nlogn)
func SmallestIntegerFinder(numbers []int) int {
	sort.Ints(numbers)
	return numbers[0]
}

//THIS FUNCTION O(n)
/*
func findSmallestInteger(arr []int) int {
    if len(arr) == 0 {
        return 0
    }

    smallest := arr[0]

    for _, num := range arr {
        if num < smallest {
            smallest = num
        }
    }

    return smallest
}

*/
