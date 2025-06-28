package main

import (
	"fmt"
	"os"
)

func main() {

	s, err := os.ReadFile("sequences/acgt.txt")
	check(err)

	//something that I should do is download the text file version of this and utilize 'os' to read that specific text file instead of just dumping it here!
	fmt.Println(ACGTCount(string(s)))

	x, err2 := os.ReadFile("sequences/transcribe.txt")
	check(err2)

	fmt.Println(Transcribe(string(x)))
	fmt.Println(ReverseComplement("AGTAAGCAGAGCCCGTAGCA"))
}

func ACGTCount(s string) (int, int, int, int) {
	var A int
	var C int
	var G int
	var T int

	for _, val := range s {
		if val == 'A' {
			A++
		}
		if val == 'C' {
			C++
		}
		if val == 'G' {
			G++
		}
		if val == 'T' {
			T++
		}
	}

	return A, C, G, T
}

func Transcribe(s string) string {
	sequence := make([]byte, len(s))
	for i, val := range s {
		if val == 'T' {
			sequence[i] = 'U'
			continue
		}
		sequence[i] = byte(val)
	}

	string_seq := string(sequence)
	return string_seq
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// re-implementation of switch case statements
func ReverseComplement(sequence string) string {
	return Reverse(Complement(sequence))
}

func Complement(sequence string) string {

	slice_seq := make([]byte, len(sequence))
	for i, val := range sequence {
		symbol := byte(val)

		switch symbol {
		case 'C':
			slice_seq[i] = 'G'
		case 'G':
			slice_seq[i] = 'C'
		case 'A':
			slice_seq[i] = 'T'
		case 'T':
			slice_seq[i] = 'A'
		}
	}

	return string(slice_seq)
}

func Reverse(sequence string) string {
	reversed := make([]byte, len(sequence))

	for i, val := range sequence {
		reversed[len(sequence)-i-1] = byte(val)
	}

	return string(reversed)
}
