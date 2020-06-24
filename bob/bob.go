// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import "strings"

// Hey should have a comment documenting it.
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
