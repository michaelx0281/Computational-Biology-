package main

import (
	"fmt"
	"io"       //needed to read from (and write to) files
	"log"      //needed for log files (errors)
	"net/http" //for accessing URLs
	//"time"
)

func main() {
	fmt.Println("Clumps.")

	/*
		text := "AAAACGTCGAAAA"
		k := 2
		L := 4
		t := 2

		fmt.Println(FindClumps(text, k, L, t))
	*/

	//rather than copying over an entire genome --> which is going to be so cluttered --> read this from a file from URL

	url := "https://bioinformaticsalgorithms.com/data/realdatasets/Replication/E_coli.txt" // accessing the url will give an response object

	resp, err := http.Get(url) //why is there an REST api cameo....scary

	//error handling
	if err != nil {
		panic(err)
	}

	//reponse was OK

	//close the connection after you're done --> saves to very end of func main
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK { // which is typically 200s
		log.Fatalf("Received non-OK status: %v", resp.Status)
	}

	genomeSymbols, err := io.ReadAll(resp.Body) //a slice of all of the symbols read from the body and an error message

	if err != nil {
		panic(err)
	}

	//we now have a slice of symbols associated with the genome]

	fmt.Println("The number of nucleotides in E. coli genome is", len(genomeSymbols))
	// //turn slice into string
	EcoliGenome := string(genomeSymbols)

	k := 9 //the significant number --> frequently
	L := 500
	t := 3

	// clumps := FindClumps(EcoliGenome, k, L, t)

	// fmt.Println("Found", len(clumps), "total patterns occuring as clumps.")

	// start := time.Now()
	// FindClumps(EcoliGenome, k, L, t)
	// elapsed := time.Since(start)
	// log.Printf("FindClumps took %s", elapsed)

	// start2 := time.Now()
	clumps2 := FindClumpsFaster(EcoliGenome, k, L, t)
	// elapsed2 := time.Since(start2)
	// log.Printf("FindClumpsFaster took %s", elapsed2)

	fmt.Println("Found", len(clumps2), "total patterns occuring as clumps.")
}

//text = "BANANASPLIT"
//first: BANANA
//second: ANANAS

//FindClumpsFaster takes as input a string text, and integers k, l, and t.
//It returns a slice of strings representing all substrings of length k that appear at least t times in a "window" of length L in the string text.

// redo this without throwing away the maps fully
func FindClumpsFaster(text string, k, L, t int) []string {
	patterns := make([]string, 0)
	n := len(text)

	//map to track whether I have added a string to patterns yet
	foundPatterns := make(map[string]bool)

	firstWindow := text[:L]

	freqMap := FrequencyTable(firstWindow, k)

	//append any patterns we find to patterns slice
	for s, freq := range freqMap {
		if freq >= t {
			patterns = append(patterns, s)
			foundPatterns[s] = true
		}
	}

	//range over all remaining possible windows of text
	for i := 1; i < n-L+1; i++ {
		//decrease by 1 the value associated with the first substring of length k in the preceding window
		oldPattern := text[i-1 : i-1+k]
		freqMap[oldPattern]--

		//clean up the map if the frequency of oldPattern was 1
		if freqMap[oldPattern] == 0 {
			delete(freqMap, oldPattern)
		}

		//add the pattern from the end of the current window
		newPattern := text[i+L-k : i+L]
		freqMap[newPattern]++

		//I have updated the frequency map :3
		//Now I need to update the patterns
		for s, freq := range freqMap {
			if freq >= t && !foundPatterns[s] {
				patterns = append(patterns, s)
				foundPatterns[s] = true
			}
		}
	}

	return patterns
}

//FindClumps takes as input a string text, and integers k, l, and t.
//It returns a slice of strings representing all substrings of length k that appear at least t times in a "window" of length L in the string text.

func FindClumps(text string, k, L, t int) []string {
	patterns := make([]string, 0)
	n := len(text)

	//map to track whether I have added a string to patterns yet
	foundPatterns := make(map[string]bool)

	//range over all possible windows of text --> the windows are of length L
	for i := 0; i < n-L+1; i++ {
		//set the current window
		window := text[i : i+L] // L is bigger than k --> within L look for frequently occuring substrings of length k in L

		//let's make the frequency table for this window
		freqMap := FrequencyTable(window, k)

		//what occurs frequently(i.e, t or more times)?
		for s, val := range freqMap {
			//append s to patterns if  s occurs frequently
			//and s doesn't already appear in patterns
			if val >= t && !foundPatterns[s] { //rather than maintaining this as a slice --> map instead --> default value of bool is false
				patterns = append(patterns, s)
				foundPatterns[s] = true
			}
		}
	}

	return patterns
}

//Contains takes as input a slice of strings and a single string s.
//It returns true if s is contained in the slice and false otherwise.

func Contains(patterns []string, s string) bool {
	for _, pattern := range patterns {
		if s == pattern {
			return true
		}
	} // after surviving the ranging --> conclude default and return false
	return false
}

//FreqencyTable takes as input a string text and an integer k.
//It returns the frequency table mapping each substring of text of legnth k to its number of occurences.

func FrequencyTable(text string, k int) map[string]int {
	freqMap := make(map[string]int) // string is the key, int is the value

	n := len(text)

	//range over all possible substrings of length k --> n - k + 1 total starting positions
	for i := 0; i < n-k+1; i++ {
		//grab current pattern
		pattern := text[i : i+k]

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
