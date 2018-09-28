// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym converts noun phrases into acronyms.
package acronym

import (
	"strings"
)

// Abbreviate shortens a noun phrase into its acronym.
func Abbreviate(s string) string {
	var acronym string
	s = strings.ToUpper(strings.Replace(s, "-", " ", -1))
	words := strings.Fields(s)
	for _, w := range words {
		acronym += string(w[0])
	}
	return strings.ToUpper(acronym)
}
