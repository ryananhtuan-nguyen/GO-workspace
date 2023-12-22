// Given a string of digits, you should replace any digit below 5 with '0' and any digit 5 and above with '1'. Return the resulting string.

// Note: input will never be an empty string

package kata

import (
	"strconv"
	"strings"
)

func FakeBin(x string) string {
	result := ""
	for _, v := range strings.Split(x, "") {
		num, _ := strconv.Atoi(v)

		if num < 5 {
			result += "0"
		} else {
			result += "1"
		}
	}
	return result

}

/*Shorter


 */
