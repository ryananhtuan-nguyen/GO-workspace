// Given a set of numbers, return the additive inverse of each. Each positive becomes negatives, and the negatives become positives.

// invert([1,2,3,4,5]) == [-1,-2,-3,-4,-5]
// invert([1,-2,3,-4,5]) == [-1,2,-3,4,-5]
// invert([]) == []

package kata

func Invert(arr []int) (result []int) {
	for _, i := range arr {
		result = append(result, -i)
	}
	return
}
