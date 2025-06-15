package main

import (
	"fmt"
	"strconv"
)
func main() {
	fmt.Println("Substrings (and subslices)!") 

	s := "Hi Lovers!"

	fmt.Println(s[1:5]) //i Lo --> up to but not including 5

	//substring at the beginning of a string ==> prefix
	//substring at the end of a string ==> suffix
	fmt.Println(s[:7]) //Hi Love --> without the front / end == infered prefix / suffix
	fmt.Println(s[4:]) //overs!

	a := make([]int, 6) //subslice
	for i := range a {
		a[i] = 2*i + 1
	}

	fmt.Println(a)
	fmt.Println(a[3:5])
	fmt.Println(a[:3])
	fmt.Println(a[4:])

	text:="abababaefwf"
	pattern:="aba"
	fmt.Println(string(strconv.Itoa(PatternCount(pattern, text))))

	fmt.Println(PatternCount(pattern, text))
	fmt.Println(StartingIndices(pattern, text))
}

//PatternCount takes as input two strings patterns and text.
//It returns the number of times that pattern occurs as a substring of text, including overlaps.

func PatternCount(pattern, text string) int { // this includes overlap!
	// substringLength := len(pattern)

	// counter := 0
	// for i := 0; i < len(text) - len(pattern) + 1; i++ {
	// 	if text[i:i+substringLength] == pattern {
	// 		counter += 1
	// 	}
	// }

	// return counter
	return len(StartingIndices(pattern, text))
}

//StartingIndices takes as input two strings pattern and text.
//It returns the collection of all starting positionas of pattern as a substring of text, including overlaps.

func StartingIndices(pattern, text string) []int {
	positions := make([]int, 0)

	n:= len(text)
	k:= len(pattern)
	
	for i:=0; i < n - k + 1; i++ {
		if pattern == text[i:i+k] {
			positions = append(positions, i)
		}
	}

	return positions
}