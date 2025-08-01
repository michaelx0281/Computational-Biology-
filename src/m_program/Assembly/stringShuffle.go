package main

import "math/rand"

// ShuffleStrings takes a collection of strings patterns as input.
// It returns a random shuffle of the strings.
func ShuffleStrings(patterns []string) []string {
	perms := rand.Perm(len(patterns))
	p := make([]string, 0)

	for _, perm := range perms {
		p = append(p, patterns[perm])
	}

	return p

}
