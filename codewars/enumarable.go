// Create a method each_cons that accepts a list and a number n, and returns cascading subsets of the list of size n, like so:

// each_cons([1,2,3,4], 2)
//   #=> [[1,2], [2,3], [3,4]]

// each_cons([1,2,3,4], 3)
//   #=> [[1,2,3],[2,3,4]]

// As you can see, the lists are cascading; ie, they overlap, but never out of order.
package enum

func EachCons(arr []int, n int) [][]int {
	var result [][]int

	if n > len(arr) || n <= 0 {
		return result
	}

	for i := 0; i <= len(arr)-n; i++ {
		result = append(result, arr[i:i+n])
	}

	return result
}
