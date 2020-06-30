//Package isogram provides isogram identification methods.
package isogram

import (
	"unicode"
)

//IsIsogram yields a boolian value indicating whether a word is an isogram.
func IsIsogram(word string) bool {
	letterCounts := make(map[rune]bool)
	for _, val := range word {
		if !unicode.IsLetter(val) {
			continue
		}
		lowerVal := unicode.ToLower(val)
		if letterCounts[lowerVal] {
			return false
		}
		letterCounts[lowerVal] = true
	}
	return true
}
