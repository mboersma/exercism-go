// Package grains calculates how many grains of wheat are on a chessboard given
// that the number on each square doubles.
package grains

import (
	"errors"
)

// Square returns two to the power of the provided integer.
func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, errors.New("value must be between 1 and 64, inclusive")
	}
	return 1 << uint64(n-1), nil
}

// Total sums the number of grains on all 64 squares of the chessboard.
func Total() uint64 {
	return 1<<64 - 1
}
