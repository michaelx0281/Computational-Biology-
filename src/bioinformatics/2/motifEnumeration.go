package main

import "github.com/michaelx0281/Computational-Biology/src/utils"

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

//this still needs more work, obviously ;-;
func MotifEnumeration(Dna []string, k, d int) []string {
	patterns := make([]string, 0)

	//range over the first only for now, then check for pattern' in each of the subsequent until the end
	for i := 0; i < len(Dna[0])-k+1; i++ {
		pattern_ := Dna[0][i : i+k]

		neighborhood := Neighbors(pattern_, d)

		if MatchPatternsSubsequent(neighborhood, Dna[1:]) {
			//add to the list
			patterns = append(patterns, pattern_)
		}
	}

	//remember to get rid of any duplicates here!
	patterns = utils.RemoveDuplicatesFromArray(patterns)

	return patterns
}

func MatchPatternsSubsequent(neighborhood, Dna []string) bool {

	k := len(neighborhood[0])

	list := make([]bool, len(Dna))

	for strand := range Dna {
		for j := 0; j < len(Dna[strand])+1; j++ {
			pattern_ := Dna[strand][strand : strand+k]

			if PatternInNeighborhood(neighborhood, pattern_) {
				list[strand] = true
			}
		}
	}

	return Tautology(list)
}

func Tautology(list []bool) bool {
	for _, boolean := range list {
		if !boolean {
			return boolean
		}
	}
	return true
}

func PatternInNeighborhood(neighborhood []string, pattern_ string) bool {
	for _, pattern := range neighborhood {
		if pattern_ == pattern {
			return true
		}
	}

	return false
}
