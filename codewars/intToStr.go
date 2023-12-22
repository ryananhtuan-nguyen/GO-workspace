// We need a function that can transform a number (integer) into a string.

// What ways of achieving this do you know?

// Examples (input --> output):
// 123  --> "123"
// 999  --> "999"
// -100 --> "-100"

package kata

import "strconv"

func NumberToString(n int) string {
	return strconv.Itoa(n)
}
