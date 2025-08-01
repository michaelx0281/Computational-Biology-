package main

import "math/rand/v2"

// SimulateReadsClean takes a genome along with a read length and a probability.
// It returns a collection of strings resulting from simulating clean reads,
// where a given position is sampled with the given probability.
func SimulateReadsClean(genome string, readLength int, probability float64) []string {

	n := len(genome)
	reads := make([]string, 0, n-readLength+1) //the third parameter would be able to pre-allocate memory for you here

	// range over all possible kmers in genome
	for i := 0; i < n-readLength+1; i++ {
		//flip a coin
		randNumber := rand.Float64()

		if randNumber < probability {
			// sample current k-mer with probability given
			currentRead := genome[i : i+readLength]
			reads = append(reads, currentRead)
		}
	}

	return reads
}
