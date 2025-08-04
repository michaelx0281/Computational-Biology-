package main

func GCSkew(genome string) int {
	var skew int

	n := len(genome)

	for i := range n + 1 { // mistake: n instead of n + 1 (range is i := 0; i < n; i++ not i <= n)
		if i == 0 {
			skew = 0
		} else {
			skew += UpdateSkew(genome, i-1)
		}
	}

	return skew
}

func GCSkewFull(genome string) []int {
	skew := make([]int, len(genome)+1)

	n := len(genome)

	for i := range n + 1 { // mistake: n instead of n + 1 (range is i := 0; i < n; i++ not i <= n)
		if i == 0 {
			skew[i] = 0
		} else {
			skew[i] = skew[i-1] + UpdateSkew(genome, i-1)
		}
	}

	return skew
}

func UpdateSkew(genome string, i int) int {
	previous := byte(genome[i])
	skew := 0

	switch previous {
	case 'G':
		skew = 1
	case 'C':
		skew = -1
	default:
		skew = 0
	}

	return skew
}
