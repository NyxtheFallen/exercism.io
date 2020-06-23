//Package proverb accepts a slice of strings and returns a proverb based on those strings.
package proverb

import "fmt"

// Proverb accepts a slice of strings and returns a proverb in a slice of equal length.
func Proverb(rhyme []string) []string {	
	lines := len(rhyme)
	if lines == 0 { return []string{} }

	proverb := make([]string, 0, lines+1)
	for i := 0; i < lines-1; i++ {
		proverb = append(proverb, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1]))
	}
	proverb = append(proverb, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))

	return proverb
}
