package main

// KmerComposition returns the k-mer composition (all k-mer substrings) of a given genome.
func KmerComposition(genome string, k int) []string {
	kmers := make([]string, 0)
	for i := 0; i < len(genome)-k+1; i++ {
		kmers = append(kmers, genome[i:i+k])
	}

	return kmers
}
