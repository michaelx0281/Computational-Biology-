package main

/*

Code Challenge: Solve the Frequent Words Problem.

Input: A string Text and an integer k.
Output: All most frequent k-mers in Text.

Currently supported languages are: Python, C++, Java, Go. We would love to add more languages in the future!

For now, to get credit for the problem, you will need to code in one of the supported languages. If you want to solve this problem in another language and are not interested in receiving points, then please check out the ungraded "Rosalind-style" problem at the end of this chapter.

*/

func FrequentWords(Text string, k int) []string {
	list := make([]string, 0)
	freqTable := make(map[string]int, 0)

	n := len(Text)

	//I can and should make each of these for loops into subroutines! (they are each self contained units!) --> name them appropiately!

	//ranges over the entire window
	for i := 0; i < n-k+1; i++ {
		window := Text[i : i+k]
		freqTable[window]++ // updates for every k-mer
	}

	maxOccurence := 0
	//range over the freqTable and add to the list the ones with the highest frequency
	for _, val := range freqTable {
		if val > maxOccurence {
			maxOccurence = val
		}
	}

	for key, val := range freqTable {
		if val == maxOccurence {
			list = append(list, key)
		}
	} // I feel like there may have been an simpler way to do this rather than two for loops, but I cannot think of the solution at the current moment.

	return list
}

//TODO --
