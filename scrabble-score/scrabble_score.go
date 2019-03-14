// Package scrabble computes the Scrabble score of a word.
package scrabble

import "strings"

// Score returns the base Scrabble score for the given word.
func Score(word string) int {
	score := 0
	for _, c := range strings.ToLower(word) {
		switch c {
		case 'a', 'e', 'i', 'o', 'u', 'l', 'n', 'r', 's', 't':
			score += 1
		case 'd', 'g':
			score += 2
		case 'b', 'c', 'm', 'p':
			score += 3
		case 'f', 'h', 'v', 'w', 'y':
			score += 4
		case 'k':
			score += 5
		case 'j', 'x':
			score += 8
		case 'q', 'z':
			score += 10
		}
	}
	return score
}
