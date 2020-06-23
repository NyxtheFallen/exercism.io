//Package acronym provides a method for converting strings to abbreviations.
package acronym

import (
	"strings"
	"unicode"
)

//Abbreviate takes a string as an argument and returns a string representing the original argument's acronym.
//Ideally, this would use a regex match, but since \b matches apostrophes, it becomes impractical.
func Abbreviate(s string) string {
	acronymSlice := strings.Split(strings.ReplaceAll(strings.ReplaceAll(s, "-", " "), "_", " "), " ")
	var firstLetters string
	for _, word := range acronymSlice {
		for _, letter := range word {
			if unicode.IsLetter(letter) {
				firstLetters += string(letter)
				break
			}
		}
	}
	acronym := strings.ToUpper(firstLetters)
	return acronym
}
