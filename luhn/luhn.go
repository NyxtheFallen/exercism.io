//Package luhn provides a method for calculating Luhn's algorithm.
package luhn

import (
	"strconv"
	"strings"
)

//Valid returns true if its input argument is a valid Luhn checksum, or false if it is not.
func Valid(num string) bool {
	num = strings.ReplaceAll(num, " ", "")

	if len(num) <= 1 {
		return false
	}

	runes := []rune(num)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	isDoubled := false
	total := 0
	for _, val := range runes {
		digit, err := strconv.Atoi(string(val))
		if err != nil {
			return false
		}
		if isDoubled {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
			total += digit
			isDoubled = !isDoubled
		} else {
			total += digit
			isDoubled = !isDoubled
		}
	}

	if total%10 != 0 {
		return false
	}
	return true
}
