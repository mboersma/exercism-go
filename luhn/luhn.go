// Package luhn determines whether or not a number is valid per the Luhn formula.
package luhn

import (
	"strings"
	"unicode"
)

// Valid returns true if a string passes the Luhn checksum algorithm.
func Valid(s string) bool {
	// Strip whitespace
	s = strings.Join(strings.Fields(s), "")

	if len(s) <= 1 {
		return false
	}

	var t = []int{}
	check := false
	for i := len(s) - 1; i >= 0; i-- {
		if !unicode.IsDigit(rune(s[i])) {
			return false
		}
		c := s[i] - '0'
		if check {
			c *= 2
			if c > 9 {
				c -= 9
			}
		}
		t = append(t, int(c))
		check = !check
	}

	sum := 0
	for _, n := range t {
		sum += n
	}

	return sum%10 == 0
}
