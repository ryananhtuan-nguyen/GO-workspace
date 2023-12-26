// Build a function that returns an array of integers from n to 1 where n>0.

// Example : n=5 --> [5,4,3,2,1]

package kata

func ReverseSeq(n int) (r []int) {
	for i := n; i > 0; i-- {
		r = append(r, i)
	}
	return
}
