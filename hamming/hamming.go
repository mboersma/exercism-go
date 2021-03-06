// Package hamming calculates the Hamming distance between two DNA strands.
package hamming

import "errors"

// Distance returns how many characters differ between two strings in the same
// locations, also known as the Hamming distance.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("strings must be the same length")
	}

	distance := 0
	for i := range a {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}
