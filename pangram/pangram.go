// Package pangram determines if a sentence uses every letter of the alphabet
// at least once.
package pangram

import "strings"

// IsPangram answers whether the given string is a pangram, based on whether or
// not it includes every letter of the alphabet at least once.
func IsPangram(input string) bool {
	seen := make(map[rune]bool)
	for _, c := range strings.ToLower(input) {
		seen[c] = true
	}
	for _, c := range "abcdefghijklmnopqrstuvwxyz" {
		if _, ok := seen[c]; !ok {
			return false
		}
	}
	return true
}
