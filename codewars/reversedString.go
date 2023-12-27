// Complete the solution so that it reverses the string passed into it.

// 'world'  =>  'dlrow'
// 'word'   =>  'drow'

package kata

func Solution(word string) string {
	rr := []rune(word)
	l := len(rr)
	for i, j := 0, l-1; i < l/2; i, j = i+1, j-1 {
		rr[i], rr[j] = rr[j], rr[i]
	}
	return string(rr)
}
