/*
package darts returns a score from 1-10, given a set of x and y coordinates.
*/
package darts

import "math"

//Return an integer score based on where a dart lands on a dartboard.
func Score (x float64, y float64) int {
	var score int
	originDistance := math.Abs(math.Hypot(x, y))

	switch {
	case originDistance <= 1: score = 10
	case originDistance <= 5: score = 5
	case originDistance <= 10: score = 1
	default: score = 0	
	}

	return score
}