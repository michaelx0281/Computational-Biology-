package main

/*
Input: Two strings
Output: The number of mismatches between strings
*/
func HammingDist(seq1, seq2 string) int {
	if len(seq1) != len(seq2) {
		panic("Unequal lengths of sequences makes it impossible to determine hamming distance.")
	}

	count := 0

	for i, char := range seq1 {
		if byte(char) != seq2[i] {
			count++
		}
	}

	return count
}
