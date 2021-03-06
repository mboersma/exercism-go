// Package series finds contiguous substrings in a string.
package series

// All returns a list of all substrings of s with length n.
func All(n int, s string) []string {
	subs := []string{}
	end := len(s) - n
	for i := 0; i <= end; i++ {
		subs = append(subs, s[i:i+n])
	}
	return subs
}

// UnsafeFirst returns the first substring of s with length n.
func UnsafeFirst(n int, s string) string {
	return s[:n]
}

// First returns the first substring of s with length n, and whether or not
// the substring requested was out of bounds.
func First(n int, s string) (first string, ok bool) {
	if n <= len(s) {
		first, ok = s[:n], true
	}
	return
}
