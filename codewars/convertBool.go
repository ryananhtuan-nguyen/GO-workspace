// Complete the method that takes a boolean value and return a "Yes" string for true, or a "No" string for false.

package kata

func BoolToWord(word bool) string {
	if word {
		return "Yes"
	}
	return "No"
}

/*
package kata

func BoolToWord(word bool) string {
  return map[bool]string{false:"No", true:"Yes"}[word]
}
*/
