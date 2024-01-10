//Given an array of integers, return a new array with each value doubled.

// For example:

// [1, 2, 3] --> [2, 4, 6]

package kata

func Maps(x []int) (res []int) {
	for _, v := range x {
		res = append(res, v*2)
	}
	return
}
