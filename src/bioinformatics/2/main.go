package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	fmt.Println("Welcome to Module 2.")
	ExerciseOne()
}

//Exercise 1

/*

Exercise Break: What is the expected number of occurrences of a 9-mer in 500 random DNA strings, each of length 1000? Assume that the sequences are formed by selecting each nucleotide (A, C, G, T) with the same probability (0.25).

Note: Express your answer as a decimal; allowable error = 0.0001.

*/

func ExerciseOne() {
	fmt.Println(SimulateNTimes(1000))
}

func SimulateNTimes(n int) float64 {
	list := make([]float64, n)

	//log each number
	for i := 0; i < n; i++ {
		list[i] = float64(SimulateOnce())
	}

	sum := SumArr(list)

	return sum / float64(n)
}

func SumArr(arr []float64) float64 {
	sum := 0.0

	for _, val := range arr {
		sum += val
	}

	return sum
}

func SimulateOnce() int {
	//Let's make a 9-mer!
	queryNineMer := GenerateKmer(9)

	count := 0
	//loop 500 times
	for i := 0; i < 500; i++ {
		dna := GenerateKmer(1000)
		count += CountOccurences(dna, queryNineMer)
	}

	//by this point, count will have the number of occurences across all of those strings
	return count
}

func CountOccurences(dna, kmer string) int {
	occurences := 0

	n := len(dna)
	k := len(kmer)

	for i := 0; i < n-k+1; i++ {
		pattern := dna[i : i+k]

		if pattern == kmer {
			occurences++
		}
	}

	return occurences
}

func GenerateKmer(k int) string {
	if k == 1 {
		return string(generate1mer())
	}

	fmt.Println("k:", k)
	sequence := make([]byte, k)
	for i := 0; i < k; i++ {

		fmt.Println("1mer:", generate1mer()) //print this out just to see
		sequence[i] = (generate1mer())[0]
	}

	return string(sequence)
}

// this is a helper function which is inaccessible outside of this file (it is private)
func generate1mer() []byte {
	n := rand.IntN(4) //4 = 4 elements, don't input 5

	var letter string

	switch n {
	case 0:
		letter = "A"
	case 1:
		letter = "T"
	case 2:
		letter = "C"
	case 3:
		letter = "G"
	}
	return []byte(letter)
}
