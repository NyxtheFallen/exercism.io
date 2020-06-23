/*
This package assists in calculating leap years!
*/
package leap

// Given an integer year, return a boolian value indicating whether or not it is a leap year.
func IsLeapYear(year int) bool {
	var leapYear bool
	remainder := year % 400

	switch {
	case remainder == 0: leapYear = true
	case remainder % 100 == 0: leapYear = false
	case remainder % 4 == 0: leapYear = true
	default: leapYear = false
	}
	return leapYear
}
