package main

// Minimizer takes a string text and an integer k as input.
// It returns the k-mer of text that is lexicographically minimum.
func Minimizer(text string, k int) string {

	windowLength := len(text) //window length, k is the length of the kmer

	minimized := text[0:k]

	//range through all possible values
	for i := 0; i < windowLength-k+1; i++ {
		currentKMer := text[i : i+k]

		if currentKMer < minimized {
			minimized = currentKMer
		}
	}

	return minimized
}
