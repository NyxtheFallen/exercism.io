//Package raindrops provides a FizzBuzz-like function.
package raindrops

import "strconv"

//Convert accepts an integer argument and returns its PlingPlangPlong value:
//Pling if the integer is divisible by 3, Plang if it is divisible by 5, Plong if 7,
//and the integer itself if none of the cases pass.
func Convert(integer int) string {
	var result string
	if integer%3 == 0 {
		result += "Pling"
	}
	if integer%5 == 0 {
		result += "Plang"
	}
	if integer%7 == 0 {
		result += "Plong"
	}
	if result == "" {
		result += strconv.Itoa(integer)
	}

	return result
}
