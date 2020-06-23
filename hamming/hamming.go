// Package hamming provides functions related to hamming distances.
package hamming

import (
	"errors"
)

// Distance calculates the Hamming Distance between two DNA strands.
func Distance(a, b string) (int, error) {
	// Make sure we compare individual utf8 runes instead of bytes
	ar, br := []rune(a), []rune(b)
	if len(ar) != len(br) {
		return 0, errors.New("dna strands of different lengths are not allowed")
	}

	var distance = 0
	for i, letter := range ar {
		if letter != br[i] {
			distance++
		}
	}
	return distance, nil
}
