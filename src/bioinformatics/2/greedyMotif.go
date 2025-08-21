package main

// import "fmt"

/*
	This is supposed to exemplify the Greedy Motif Search.

	Profile-most Probable k-mer Problem: Find a Profile-most probable k-mer in a string.

	Input: A string Text, an integer k, and a 4 × k matrix Profile.
	Output: A Profile-most probable k-mer in Text.
*/

func ProfileMostProbableKmer(Text string, k int, Profile [4][]float64) string {
	//k represents the length of the kmer taken from text

	n := len(Text)

	maxProbability := 0.0
	motif := ""

	for i := 0; i < n-k+1; i++ {
		//now for each of these possibilities, you want to maximize the score of multiplying all of these different profiles together

		//ACGT in terms of order in the profile
		pattern := Text[i : i+k]

		distribution := make([]float64, k)

		for j, symbol := range pattern {
			s := byte(symbol)
			switch s {
			case 'A':
				distribution[j] = Profile[0][j]
			case 'C':
				distribution[j] = Profile[1][j]
			case 'G':
				distribution[j] = Profile[2][j]
			case 'T':
				distribution[j] = Profile[3][j]
			}
		}

		product := calcProbability(distribution)

		// fmt.Println("Product", i, product)

		if product > maxProbability {
			maxProbability = product
			// fmt.Println(maxProbability)
			motif = pattern

			// fmt.Println("Motif:", motif, "%", maxProbability)
		}
	}

	return motif
}

func calcProbability(distribution []float64) float64 {
	product := 1.0

	for _, float := range distribution {
		product *= float
	}

	return product
}

/*
Code Challenge: Implement GreedyMotifSearch().

Input: Integers k and t, followed by a space-separated collection of strings Dna.
Output: A collection of strings BestMotifs resulting from applying GreedyMotifSearch(Dna, k, t). If at any step you find more than one Profile-most probable k-mer in a given string, use the one occurring first.

Pseudocode:

GreedyMotifSearch(Dna, k, t)
    BestMotifs ← motif matrix formed by first k-mers in each string from Dna
    for each k-mer Motif in the first string from Dna
        Motif1 ← Motif
        for i = 2 to t
            form Profile from motifs Motif1, …, Motifi - 1
            Motifi ← Profile-most probable k-mer in the i-th string in Dna
        Motifs ← (Motif1, …, Motift)
        if Score(Motifs) < Score(BestMotifs)
            BestMotifs ← Motifs
    return BestMotifs
*/
