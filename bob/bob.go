/*
Package bob is just a silly response generator.
*/
package bob

import "strings"

// Hey generates a limited number of responses based on an input remark.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	isYelling := strings.ToUpper(remark) == remark && strings.ToLower(remark) != remark //filter for numerics
	switch {
	case strings.HasSuffix(remark, "?"):
		switch {
		case isYelling:
			return "Calm down, I know what I'm doing!"
		default:
			return "Sure."
		}
	case isYelling:
		return "Whoa, chill out!"
	case remark == "":
		return "Fine. Be that way!"
	default:
		return "Whatever."
	}
}
