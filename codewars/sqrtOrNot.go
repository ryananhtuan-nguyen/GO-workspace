// Write a method, that will get an integer array as parameter and will process every number from this array.

// Return a new array with processing every number of the input-array like this:

// If the number has an integer square root, take this, otherwise square the number.

// Example
// [4,3,9,7,2,1] -> [2,9,3,49,4,1]
// Notes
// The input array will always contain only positive numbers, and will never be empty or null.

package kata

import "math"

func SquareOrSquareRoot(arr []int) []int {
	results := make([]int, len(arr))

	for i, x := range arr {
		sqrt := math.Sqrt(float64(x))

		if sqrt == math.Floor(sqrt) {
			results[i] = int(sqrt)
		} else {
			results[i] = x * x
		}
	}

	return results
}

//NOTE:
// It appears that the issue lies in the comparison of floating-point numbers. When dealing with floating-point arithmetic, small rounding errors may occur, leading to differences in the expected and actual results.
//Another way to do this
/*

package kata

import "math"

const epsilon = 1e-9

func SquareOrSquareRoot(arr []int) []int {
	var result []int

	if arr == nil {
		return result
	}

	for _, v := range arr {
		squareRoot := math.Sqrt(float64(v))
		if math.Abs(squareRoot-math.Trunc(squareRoot)) < epsilon {
			result = append(result, int(squareRoot))
		} else {
			result = append(result, v*v)
		}
	}

	return result
}

*/
