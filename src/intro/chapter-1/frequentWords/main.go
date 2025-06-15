package main

import (
	"fmt"
)

func main() {
	fmt.Println("Finding frequent words.")

	//go back and do some review on maps later / soon

	// this is the short region of genome that is known experimentally for being important for vibrio cholerae in DNA replication
	text := "ATCAATGATCAACGTAAGCTTCTAAGCATGATCAAGGTGCTCACACAGTTTATCCACAACCTGAGTGGATGACATCAAGATAGGTCGTTGTATCTCCTTCCTCTCGTACTCTCATGACCACGGAAAGATGATCAAGAGAGGATGATTTCTTGGCCATATCGCAATGAATACTTGTGACTTGTGCTTCCAATTGACATCTTCAGCGCCATATTGCGCTGGCCAAGGTGACGGAGCGGGATTACGAAAGCATGATCATGGCTGTTGTTCTGTTTATCTTGTTTTGACTGAGACTTGTTAGGATAGACGGTTTTTCATCACTGACTAGCCAAAGCCTTACTCTGCCTGACATCGACCGTAAATTGATAATGAATTTACATGCTTCCGCGACGATTTACCTCTTGATCATCGATCCGATTGAAGATCTTCAATTGTTAATTCTCTTGCCTCGACTCATAGCCATGATGAGCTCTTGATCATGTTTCCTTAACCCTCTATTTTTTACGGAAGAATGATCAAGCTGCTGCTCTTGATCATCGTTTC"


	/* 
		Finding frequent words.
		Frequent pattern(s) found when k is 9 are [CTCTTGATC *ATGATCAAG TCTTGATCA *CTTGATCAT]
		The maximum number of occurrences is 3

		The starred ones are reverse complements! (so this message is printed 3 times on each pairing strand of DNA!)
	*/
	k := 9 //the biologically interesting value of k = 9

	freqMap := FrequencyTable(text, k)

	freqPatterns := FindFrequentWords(text, k)

	fmt.Println("Frequent pattern(s) found when k is", k, "are", freqPatterns)

	fmt.Println("The maximum number of occurrences is", freqMap[freqPatterns[0]])

	// fmt.Println(freqMap["TGA"]) <-- badd this is hard coding. need to soft code it
}

//FindFrequentWords takes as input a string text and an integer k.
//It returns a slice of strings corresponding to the substring(s) of length k that occur most frequently in text.
func FindFrequentWords(text string, k int) []string {
	freqPatterns := make([]string, 0)


	freqMap := FrequencyTable(text, k)
	//once I have the frequency table, find the maximum value --> range over the map
	max := MaxMapValue(freqMap)

	//what keys of this table reach the max value?
	
	//range over the frequencyMap
	for pattern, val := range freqMap {
		if val == max {
			//append!
			freqPatterns = append(freqPatterns, pattern)
		}
	}

	return freqPatterns
}

//MaxMapValue takes as input a map of strings to integers.
//It returns the maximum value found in the map.
func MaxMapValue(dict map[string]int) int {
	m := 0
	firstTimeThrough := true

	//range over dict/map adn update m every time I find a value that is larger
	for _, val := range dict {
		if firstTimeThrough || val > m {
			m = val
			firstTimeThrough = false //as soon as you encounter the first element --> turn it off 
		}
	}

	return m
}

//FreqencyTable takes as input a string text and an integer k.
//It returns the frequency table mapping each substring of text of legnth k to its number of occurences.

func FrequencyTable(text string, k int) map[string]int {
	freqMap := make(map[string]int) // string is the key, int is the value

	n:= len(text)


	//range over all possible substrings of length k --> n - k + 1 total starting positions
	for i := 0; i < n - k + 1; i++ {
		//grab current pattern
		pattern := text[i:i+k]

		//updating the value of the freqMap associated with pattern

		//exists --> add one to it

		// _, exists := freqMap[pattern] // exists is a boolean value that could also be returned by the freqMap

		// if !exists {
		// 	//pattern has not been encounter4ed
		// 	freqMap[pattern] = 1
		// } else [
		// 	//pattern has been encountered +=1
		// 	freqMap[pattern]++
		// ]

		// x := freqMap[pattern] // x would would get the default value of keys --> it would get 0

		freqMap[pattern]++ // this is usuable only because the default value of key ints is 0
		// if pattern is a key, this is what we want
		//if it's nt, freqMap[pattern] gets created,
		//gets set equal to zero, then incremented


	}


	return freqMap
}