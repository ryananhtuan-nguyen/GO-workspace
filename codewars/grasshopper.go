// Write a program that finds the summation of every number from 1 to num. The number will always be a positive integer greater than 0. Your function only needs to return the result, what is shown between parentheses in the example below is how you reach that result and it's not part of it, see the sample tests.

package kata

func Summation(n int) int {
	return n * (n + 1) / 2
}

/*
====================Traditional For loop:

func Summation(n int) int {
    // the sleeper must awaken!

  count := 0

  for i := 1; i <= n; i++{
    count += i

  }
  return count
}

*/
