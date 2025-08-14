package main

import (
	"fmt"

	"github.com/michaelx0281/Computational-Biology/src/utils"
)

/*

MotifEnumeration(Dna, k, d)
    Patterns ← an empty set
    for each k-mer Pattern in Dna
        for each k-mer Pattern’ differing from Pattern by at most d mismatches
            if Pattern' appears in each string from Dna with at most d mismatches
                add Pattern' to Patterns
    remove duplicates from Patterns
    return Patterns

*/

//NOTE: Only MotifEnumeration2 works here! The first one is slightly different! (curse of hyper-modularity :c))

// this still needs more work, obviously ;-;
func MotifEnumeration(Dna []string, k, d int) []string {
	patterns := make([]string, 0)

	//range over the first only for now, then check for pattern' in each of the subsequent until the end
	for i := 0; i < len(Dna[0])-k+1; i++ {
		pattern_ := Dna[0][i : i+k]

		neighborhood := Neighbors(pattern_, d)
		fmt.Println(neighborhood)

		//nothing is getting appened after this block so far
		inNeighborhood, extras := MatchPatternsSubsequent(neighborhood, Dna[1:], d)
		if inNeighborhood {
			//add to the list
			patterns = append(patterns, pattern_)
		}

		if len(extras) > 0 {
			patterns = append(patterns, extras...)
		}
	}

	//remember to get rid of any duplicates here!
	patterns = utils.RemoveDuplicatesFromArray(patterns)

	return patterns
}

func MatchPatternsSubsequent(neighborhood, Dna []string, d int) (bool, []string) {

	k := len(neighborhood[0])

	list := make([]bool, len(Dna))
	extras := make([]string, 0)

	for strand := range Dna {
		for j := 0; j < len(Dna[strand])-k+1; j++ {
			pattern_ := (Dna[strand])[j : j+k] //made the mistake here of [strand : strand + k] instead of using j

			primeNeigbors := Neighbors(pattern_, d)

			for _, pattern_ := range primeNeigbors {
				if PatternInNeighborhood(neighborhood, pattern_, d) {
					list[strand] = true
					extras = append(extras, pattern_)
				}
			}
		}
	}

	return Tautology(list), extras
}

// the first one has evolved to be unmanageable
func MotifEnumeration2(Dna []string, k, d int) []string {
	motifs := make([]string, 0)

	hash := make(map[string]int)
	countOnce := make(map[string]bool)

	neighbors := MotifNeighbors(Dna[0], k, d)
	//now check if this would be a match with the rest of the strings of Dna
	for i := 1; i < len(Dna); i++ {
		for j := 0; j < len(Dna[i])-k+1; j++ {
			pattern := Dna[i][j : j+k]

			for _, n := range neighbors {
				if HammingDist(n, pattern) <= d && countOnce[n] == false {
					//update some sort of a tracker
					hash[n]++
					countOnce[n] = true
				}
			}
		}

		//reset all of the booleans in countOnce()
		for n := range countOnce {
			countOnce[n] = false
		}
	}

	// fmt.Println(hash)

	for motif, val := range hash {
		if val == len(Dna)-1 {
			motifs = append(motifs, motif)
		}
	}

	utils.RemoveDuplicatesFromArray(motifs)

	return motifs
}

func MotifNeighbors(Text string, k, d int) []string {
	list := make([]string, 0)

	n := len(Text)

	for i := 0; i < n-k+1; i++ {
		pattern := Text[i : i+k]

		neighbors := Neighbors(pattern, d)

		for j := 0; j < len(neighbors); j++ {
			list = append(list, neighbors[j])
		}
	}

	list = utils.RemoveDuplicatesFromArray(list)

	return list
}

func Tautology(list []bool) bool {
	for _, boolean := range list {
		if !boolean {
			return boolean
		}
	}
	return true
}

func PatternInNeighborhood(neighborhood []string, pattern_ string, d int) bool {

	for _, pattern := range neighborhood { //changed this up but it looks uglier now!
		if pattern_ == pattern {
			return true
		}
		// if HammingDist(pattern_, pattern) == d {
		// 	return true
		// }
	}

	return false
}
