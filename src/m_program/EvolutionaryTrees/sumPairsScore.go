package main

import "fmt"

// SumPairsScore takes two multiple alignments as well as two indices, and scoring
// parameters.
// It returns the sum of pairs score of the corresponding columns
// in the two alignments, using the specified scoring parameters.
func SumPairsScore(align1 Alignment, align2 Alignment,
	idx1 int, idx2 int, match float64, mismatch float64, gap float64) float64 {

	//NOTES: Aligment => []string --> iterate through each of them and compare to arrive at the scores
	// No need to iterate through all of it --> just check the specific indices

	//Here are the corresponding scores below
	//align1 <-> idx1
	//align2 <-> idx2
	matchCount := 0
	mismatchCount := 0
	gapCount := 0

	//Loop through align1
	//@index 1
	for _, char1 := range align1[idx1] {
		for _, char2 := range align2[idx2] {
			//in c
			if Gap(byte(char1), byte(char2)) {
				gapCount++
			}
			if char1 == char2 && char1 != '-' && char2 != '-' {
				matchCount++
			}
			if char1 != char2 && char1 != '-' && char2 != '-' {
				mismatchCount++
			}
		}
	}
	//Loop through align2
	//@index2
	fmt.Println(gapCount, mismatchCount, matchCount)
	return -gap*float64(gapCount) - mismatch*float64(mismatchCount) + match*float64(matchCount)

}

func Gap(char1, char2 byte) bool {
	return (char1 == '-' && char2 != '-') || (char2 == '-' && char1 != '-')
}
