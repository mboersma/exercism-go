// Package bob responds to conversation as a lackadaisical teenager would.
package bob

import (
	"strings"
	"unicode"
)

// HasLetter returns whether or not a string contains an alphabetic character.
func HasLetter(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

// Hey returns a lackadaisical teenager's response to the remark provided.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	if remark == "" {
		return "Fine. Be that way!"
	}

	yell := HasLetter(remark) && strings.ToUpper(remark) == remark
	question := strings.HasSuffix(remark, "?")

	if yell && question {
		return "Calm down, I know what I'm doing!"
	}
	if question {
		return "Sure."
	}
	if yell {
		return "Whoa, chill out!"
	}
	return "Whatever."
}
