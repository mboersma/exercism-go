// Package darts scores tosses in a game of Darts.
package darts

import "math"

// Score returns the points earned for a toss of a dart at the target.
func Score(x, y float64) int {
	// The distance formula is √((x₁-x₂)²+(y₁-y₂)²), but x₂ and y₂ are 0
	// at the origin, so it simplifies to √(x₁²+y₁²).
	distance := math.Sqrt(x*x + y*y)
	switch {
	case distance <= 1: // inner circle!
		return 10
	case distance <= 5: // middle circle
		return 5
	case distance <= 10: // outer circle
		return 1
	default: // outside the target...
		return 0
	}
}
