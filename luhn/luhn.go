//Package luhn provides a method for calculating Luhn's algorithm.
package luhn

import (
	"strings"
	"unicode"
)

//Valid returns true if its input argument is a valid Luhn checksum, or false if it is not.
func Valid(num string) bool {
	num = strings.ReplaceAll(num, " ", "")

	if len(num) <= 1 {
		return false
	}

	runes := []rune(num)

	isDoubled := false
	total := 0
	for i := len(runes) - 1; i >= 0; i-- {
		if !unicode.IsDigit(runes[i]) {
			return false
		}
		digit := int(runes[i] - '0')
		if isDoubled {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		total += digit
		isDoubled = !isDoubled
	}

	if total%10 != 0 {
		return false
	}

	return true
}
