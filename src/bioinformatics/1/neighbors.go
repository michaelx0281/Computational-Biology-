package main

import "github.com/michaelx0281/Computational-Biology/src/utils"

//Given a string, return the neighboor of it!
func Neigbors(Pattern string, d int) []string {
	neigborhood := make([]string, 0)
	//There are two functions!
	s := SuffixRandomizedNeighbors(Pattern, d)
	p := PrefixRandomizedNeighbors(Pattern) //don't need to check d for this one. It will return everything hamming dist 1.

	//there needs to be more of a fix before this is done
	return utils.RemoveDuplicatesFromArray(s + p)
}

func ReplaceLetter(pattern []byte, letter byte, index int) []byte {
	pattern[index] = letter
	return pattern
}

//make this part a recursive function? //okay, this pattern is clearly not working right now / at this moment --> use the pseudocode from the website ot actually recurse instead of trying to do it dynamically
func SuffixRandomizedNeighbors(Pattern string, d int) []string {
	//convert pattern into list of bytes
	pattern := make([]byte, len(Pattern))
	neigbors := make([]string, 0)

	for i, letter := range Pattern {
		pattern[i] = byte(letter)
	}

	for i := 1; i < len(pattern); i++ {
		letter := pattern[i]

		stringedLetter := string(letter)

		nucleotides := Nucleotides()

		for i, val := range nucleotides {
			if stringedLetter == val {
				nucleotides = nucleotides[:1] + nucleotides[i+1:]
			}
		}

	}

	return []string{}
}

//This one is the simplest one! Let's do it first!
func PrefixRandomizedNeighbors(Pattern string) []string {
	neigbors := make([]string, 4)

	suffix := Suffix(Pattern)
	//Generate random prefixes!
	prefixes := Nucleotides()

	for i := 0; i < 4; i++ {
		neigbors[i] = prefixes[i] + suffix
	}

	return neigbors
}

func Suffix(Pattern string) string {
	return Pattern[1:]
}

func Nucleotides() []string {
	return []string{"A", "T", "C", "G"}
}
