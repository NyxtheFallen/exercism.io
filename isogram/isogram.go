//Package isogram provides isogram identification methods.
package isogram

import (
	"unicode"
)

//IsIsogram yields a boolian value indicating whether a word is an isogram.
func IsIsogram(word string) bool {
	letters := []rune(word)
	letterCounts := make(map[rune]int)
	for _, val := range letters {
		if unicode.IsLetter(val) {
			lowerVal := unicode.ToLower(val)
			letterCounts[lowerVal]++
			if letterCounts[lowerVal] > 1 {
				return false
			}
		}
	}
	return true
}
