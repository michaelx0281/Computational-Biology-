package main

/*
	Here is the pseudocode!

	(Autograded) Code Challenge: Implement MedianString().

	Input: An integer k, followed by a space-separated collection of strings Dna.
	Output: A k-mer Pattern that minimizes d(Pattern, Dna) among all possible choices of k-mers. (If there are multiple such strings Pattern, then you may return any one.)

	MedianString(Dna, k)
		distance ← ∞
		for each k-mer Pattern from AA…AA to TT…TT
			if distance > d(Pattern, Dna)
				distance ← d(Pattern, Dna)
				Median ← Pattern
		return Median
*/

func MedianString(k int, Dna []string) string {
	Median := ""
	distance := 100000000

	//generate all of the k-mers first!
	kmers := Neighbors(GenerateSeed(k), k)

	for _, kmer := range kmers {
		d := dSummation(kmer, Dna)
		if distance > d {
			distance = d
			Median = kmer
		}
	}
	return Median
}

func GenerateSeed(length int) string {
	s := make([]byte, length)

	for i := 0; i < length; i++ {
		s[i] = 'A'
	}

	return string(s)
}

/*
	MedianString(Dna, k)
		distance ← ∞
		for each k-mer Pattern from AA…AA to TT…TT
			if distance > d(Pattern, Dna)
				distance ← d(Pattern, Dna)
				Median ← Pattern
		return Median
*/

// func Kmers(k int, Text string) []string {
// 	n := len(Text)

// 	kmers := make([]string, n-k) //this in length starting at 1

// 	for i := 0; i < n-k+1; i++ {
// 		pattern := Text[i : i+k]

// 		kmers[i] = pattern
// 	}

// 	return kmers
// }

// func d(Pattern string, Dna []string) int {
// 	d := 0
// 	for _, strand := range Dna {
// 		d += HammingDist(strand, Pattern)
// 	}

// 	return d
// }

func dSummation(Pattern string, Dna []string) int {
	sum := 0

	for strand := range Dna {
		sum += d(Pattern, Dna[strand])
	}

	return sum

}

func d(Pattern, Text string) int {

	n := len(Text)
	k := len(Pattern)

	distances := make([]int, n-k+1)

	for i := 0; i < n-k+1; i++ {
		pattern := Text[i : i+k]

		distances[i] = HammingDist(Pattern, pattern)
	}

	return MinInt(distances)
}

func Motif(Pattern, Text string, d int) string { //slightly redudant, maybe optimize some other time?
	n := len(Text)
	k := len(Pattern)

	for i := 0; i < n-k+1; i++ {
		pattern := Text[i : i+k]

		if HammingDist(pattern, Pattern) == d {
			return pattern
		}
	}
	return Pattern //this indicates an error or d of less than 1
}

func MinInt(list []int) int {

	min := 1000000000

	for _, value := range list {
		if value < min {
			min = value
		}
	}

	return min
}
