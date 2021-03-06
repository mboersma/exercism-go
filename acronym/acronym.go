// Package acronym converts noun phrases into acronyms.
package acronym

import (
	"strings"
)

// Abbreviate shortens a noun phrase into its acronym.
func Abbreviate(s string) string {
	var acronym string
	s = strings.Replace(strings.Replace(s, "_", " ", -1), "-", " ", -1)
	words := strings.Fields(s)
	for _, w := range words {
		acronym += string(w[0])
	}
	return strings.ToUpper(acronym)
}
