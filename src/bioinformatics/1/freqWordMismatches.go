package main

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

func FrequentWordsWithMistmatches(Text string, k int, d int) []string {
	patterns := make([]string, 0)
	freqMap := make(map[string]int)

	n := len(Text)
	for i := 0; i < n-k+1; i++ { //think about the way that this could be optimized later!
		pattern := Text[i : i+k]
		neighborhood := Neighbors(pattern, d) //generating a bunch of different possible values
		for j := range neighborhood {
			neigbor := neighborhood[j] //this is the current word neighbor
			freqMap[neigbor]++
		}
		m := MaxMap(freqMap)

		//make into subroutine with a good name
		for pattern, val := range freqMap {
			if val == m {
				patterns = append(patterns, pattern)
			}
		}
	}

	return patterns
}

//Returns the largest value found within the hashmap!
func MaxMap(freqMap map[string]int) int {
	max := 0

	for _, val := range freqMap {
		if val > max {
			max = val
		}
	}

	return max
}
