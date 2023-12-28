// Create a function that returns the CSV representation of a two-dimensional numeric array.

// Example:

// input:
//    [[ 0, 1, 2, 3, 4 ],
//     [ 10,11,12,13,14 ],
//     [ 20,21,22,23,24 ],
//     [ 30,31,32,33,34 ]]

// output:
//
//	 '0,1,2,3,4\n'
//	+'10,11,12,13,14\n'
//	+'20,21,22,23,24\n'
//	+'30,31,32,33,34'
//
// Array's length > 2.
package kata

import (
	"fmt"
	"strings"
)

//my version

func ToCsvText(array [][]int) string {
	// your code here
	var r []string
	for _, v := range array {
		r = append(r, strings.Trim(strings.Join(strings.Fields(fmt.Sprint(v)), ","), "[]"))
	}
	return strings.Join(r, "\n")
}

/*More understandable ver
package kata

import ("strings"; "fmt")

func ToCsvText(array [][]int) string {
  array2 := []string{}
  for _ , a := range array {
    k := []string{}
    for _, n := range a {
      k = append(k, fmt.Sprint(n))
    }
    array2 = append(array2, strings.Join(k,","))
  }
  return strings.Join(array2, "\n")
}
*/
