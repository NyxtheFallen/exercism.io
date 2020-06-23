/*
Package etl provides methods for taking old-format data and transforming it into new-format data.
*/
package etl

import "strings"

//Transform receives old-format data and converts it to new-format data.
func Transform(legacy map[int][]string) map[string]int {
	var newFormat = make(map[string]int)
	for key, list := range legacy {
		for _, letter := range list {
			newFormat[strings.ToLower(letter)] = key
		}
	}
	return newFormat
}
