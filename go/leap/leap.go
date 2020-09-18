// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package leap : solution for exercise leap.
package leap

// IsLeapYear report if the input year is a leap year.
func IsLeapYear(year int) bool {
	if year%100 == 0 {
		if year%400 == 0 {
			return true
		}
	} else {
		if year%4 == 0 {
			return true
		}
	}
	return false
}
