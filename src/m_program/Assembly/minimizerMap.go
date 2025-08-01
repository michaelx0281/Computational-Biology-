package main

import "slices"

// StringIndex is a type that will map a minimizer string to its list of indices
// in a read set corresponding to reads with this minimizer.
type StringIndex map[string][]int

// BuildMinimizerMap takes a collection of reads, integer k and integer windowLength.
// It returns a map of k-mers to the indices of the reads in the list having this k-mer minimizer.
func BuildMinimizerMap(reads []string, k int, windowLength int) StringIndex {

	dict := make(StringIndex)

	for index, read := range reads {

		minimized := make([]string, 0)
		for i := 0; i < len(read)-windowLength+1; i++ {
			if !slices.Contains(minimized, Minimizer(read[i:i+windowLength], k)) {
				minimized = append(minimized, Minimizer(read[i:i+windowLength], k))
			}

		}

		for _, read := range minimized {
			dict[read] = append(dict[read], index)
		}
	}

	return dict
}
