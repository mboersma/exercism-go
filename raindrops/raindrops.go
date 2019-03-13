// Package raindrops prints certain factors of a number as raindrop sounds.
package raindrops

import "strconv"

// Convert prints "Pling," "Plang," and "Plong" when a number has factors of
// 3, 5, and 7 respectively. Otherwise, it prints the original number.
func Convert(input int) string {
	output := ""
	if input%3 == 0 {
		output += "Pling"
	}
	if input%5 == 0 {
		output += "Plang"
	}
	if input%7 == 0 {
		output += "Plong"
	}
	if len(output) == 0 {
		output = strconv.Itoa(input)
	}
	return output
}
