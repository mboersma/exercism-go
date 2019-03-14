// Package diffsquares computes the difference between the square of the sum and the
// sum of the squares of the first N natural numbers.
package diffsquares

// SquareOfSum returns the square of the sum of the first n natural numbers.
func SquareOfSum(n int) int {
	// Summation formula from https://en.wikipedia.org/wiki/1_%2B_2_%2B_3_%2B_4_%2B_%E2%8B%AF#Partial_sums
	sum := (n * (n + 1)) / 2
	return sum * sum
}

// SumOfSquares returns the sum of the squares of the first n natural numbers.
func SumOfSquares(n int) int {
	// Summation formula from https://brilliant.org/wiki/sum-of-n-n2-or-n3/
	return (n * (n + 1) * (2*n + 1)) / 6
}

// Difference returns the difference between the square of the sum and the sum
// of the squares of the first n natural numbers.
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
