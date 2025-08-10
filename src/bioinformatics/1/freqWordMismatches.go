package main

import (
	// "fmt"

	"github.com/michaelx0281/Computational-Biology/src/utils"
)

/*
Here's the Pseudocode:

FrequentWordsWithMismatches(Text, k, d)
    Patterns ← an array of strings of length 0
    freqMap ← empty map
    n ← |Text|
    for i ← 0 to n - k
        Pattern ← Text(i, k)
        neighborhood ← Neighbors(Pattern, d)
        for j ← 0 to |neighborhood| - 1
            neighbor ← neighborhood[j]
            if freqMap[neighbor] doesn't exist
                freqMap[neighbor] ← 1
            else
                freqMap[neighbor] ← freqMap[neighbor] + 1
    m ← MaxMap(freqMap)
    for every key Pattern in freqMap
        if freqMap[Pattern] = m
            append Pattern to Patterns
    return Patterns
*/

//This is a good use-case example of top-down programming.
//The problems seems very complicated at first, but most of it could be broken down into smaller chunks--> just keep in mind the original intent of the problem in the back of your head
//That way, you are not relying only on the given pseudocode templates to work on this problem!

func FrequentWordsWithMismatches(Text string, k int, d int) []string {
	patterns := make([]string, 0)
	freqMap := make(map[string]int)

	n := len(Text)
	for i := 0; i < n-k+1; i++ { //think about the way that this could be optimized later!
		pattern := Text[i : i+k]
		neighborhood := Neighbors(pattern, d) //generating a bunch of different possible values
		for j := range neighborhood {
			neighbor := neighborhood[j] //this is the current word neighbor
			freqMap[neighbor]++
		}
	}

	m := MaxMap(freqMap)

	//make into subroutine with a good name
	for pattern, val := range freqMap {
		if val == m {
			patterns = append(patterns, pattern)
		}
	}

	return patterns
}

// Returns the largest value found within the hashmap!
func MaxMap(freqMap map[string]int) int {
	max := 0

	for _, val := range freqMap {
		if val > max {
			max = val
		}
	}

	return max
}

func CountD(Pattern, Text string, d int) int {
	return len(ApproxMatching(Text, Pattern, d))
}

/*

For this next function, we would like to maximize the value of CountD(Pattern, Text, d) + CountD(Patternrc, Text, d) of all k-mers Pattern (yes the same Pattern in these two parameters!)

*/

func FrequentWordsMismatchesReverseComplements(Text string, k int, d int) []string {

	patterns := make([]string, 0)
	freqMap := make(map[string]int)

	n := len(Text)
	for i := 0; i < n-k+1; i++ { //think about the way that this could be optimized later!
		pattern := Text[i : i+k]
		neighborhood := Neighbors(pattern, d) //generating a bunch of different possible values
		for j := range neighborhood {
			neighbor := neighborhood[j] //this is the current word neighbor
			freqMap[neighbor]++
		}
	}

	sumMap := make(map[string]int)

	//make into subroutine with a good name
	for pattern := range freqMap {
		sum := CountD(pattern, Text, d) +
			CountD(ReverseComplement(pattern), Text, d)

		//now lets store each sum into a map, with pattern as the key

		sumMap[pattern] = sum
	}

	m := MaxMap(sumMap)

	for pattern, sum := range sumMap {
		if m == sum {
			patterns = append(patterns, pattern)
			patterns = append(patterns, ReverseComplement(pattern))
		}
	}

	patterns = utils.RemoveDuplicatesFromArray(patterns)

	return patterns
}
