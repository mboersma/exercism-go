// Package collatzconjecture processes numbers according to the
// "Collatz Conjecture," or 3x+1 problem.
package collatzconjecture

import "errors"

// CollatzConjecture returns the number of steps required to reach 1 from a
// given input when processed by the Collatz Conjecture rules.
func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return -1, errors.New("input must be positive")
	}
	steps := 0
	for n > 1 {
		steps++
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
	}
	return steps, nil
}
