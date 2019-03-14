// Package isogram determines if a word or phrase is a nonpattern word, or isogram.
package isogram

import "strings"

// IsIsogram returns true if the given string doesn't contain a repeating letter.
func IsIsogram(s string) bool {
	found := make(map[rune]bool)
	for _, c := range strings.ToLower(s) {
		if _, ok := found[c]; ok {
			return false
		}
		if c >= 'a' && c <= 'z' {
			found[c] = true
		}
	}
	return true
}
