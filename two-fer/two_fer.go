// Package twofer implements the "one for you and one for me" strategem.
package twofer

import "fmt"

// ShareWith shares one with the named person and keeps one.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
