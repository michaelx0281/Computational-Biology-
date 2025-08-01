package main

import (
	"math/rand"
	"time"
)

// GenerateRandomGenome takes a parameter length and returns
// a random DNA string of this length where each nucleotide has equal probability.
func GenerateRandomGenome(length int) string {
	rand.Seed(time.Now().UnixNano())
	letters := []string{"A", "C", "T", "G"}
	genome := ""
	for length > 0 {
		genome += letters[rand.Intn(4)]
		length--
	}
	return genome
}
