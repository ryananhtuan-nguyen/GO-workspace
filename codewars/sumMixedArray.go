// Given an array of integers as strings and numbers, return the sum of the array values as if all were numbers.

// Return your answer as a number.

package kata

import "strconv"

func SumMix(arr []any) int {

	sum := 0

	for _, val := range arr {

		switch val := val.(type) {
		case int:
			sum += val
		case string:
			k, _ := strconv.Atoi(val)
			sum += k
		}

	}

	return sum
}
