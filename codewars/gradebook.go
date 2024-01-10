// Grade book
// Complete the function so that it finds the average of the three scores passed to it and returns the letter value associated with that grade.

// Numerical Score	Letter Grade
// 90 <= score <= 100	'A'
// 80 <= score < 90	'B'
// 70 <= score < 80	'C'
// 60 <= score < 70	'D'
// 0 <= score < 60	'F'
// Tested values are all between 0 and 100. Theres is no need to check for negative values or values greater than 100.

package kata

func GetGrade(a, b, c int) rune {
	average := (a + b + c) / 3
	switch {
	case average >= 90 && average <= 100:
		return 'A'
	case average >= 80 && average < 90:
		return 'B'
	case average >= 70 && average < 80:
		return 'C'
	case average >= 60 && average < 70:
		return 'D'
	default:
		return 'F'
	}
}
