package main

// ProgressiveAlign takes two (multiple) alignments as input and
// returns a multiple alignment corresponding to combining the two
// alignments according to the Clustal dynamic programming heuristic.
func ProgressiveAlign(align1 Alignment, align2 Alignment,
	match float64, mismatch float64, gap float64, supergap float64) Alignment {

	// get the score table
	scoreTable := ProgressiveAlignmentScoreTable(align1, align2, match, mismatch, gap, supergap)

	// this is just like global alignment: start at bottom right and backtrack, building the alignment as we go

	row, col := len(align1[0]), len(align2[0])

	numStrings1, numStrings2 := len(align1), len(align2)

	// new alignment has total number of strings equal to # of rows of alignment 1 + # rows of alignment 2
	alignment := make(Alignment, numStrings1+numStrings2)

	// we want to be careful of string conctenations, so we make a slice of bytes
	allRows := make([][]byte, numStrings1+numStrings2)

	// range over rows and make them
	for i := range allRows {
		allRows[i] = make([]byte, 0)
	}

	for row > 0 && col > 0 {
		// range over the interior of the matrix
		// check your three values
		if scoreTable[row][col] == scoreTable[row][col-1]-supergap {
			//LEFT
			//alignment1 gets a bunch of gaps
			for i := 0; i < numStrings1; i++ {
				allRows[i] = append([]byte{'-'}, allRows[i]...)
			}
			//alignment 2 gets current column of alignment
			for j := 0; j < numStrings2; j++ {
				allRows[numStrings1+j] = append([]byte{align2[j][col-1]}, allRows[numStrings1+j]...)
			}
			col--
		} else if scoreTable[row][col] == scoreTable[row-1][col] {
			//UP

			//append
			for j := 0; j < numStrings1; j++ {
				allRows[j] = append([]byte{align1[j][row-1]}, allRows[j]...)
			}

			//bunch of gaps in alignment 2
			for i := 0; i < numStrings1; i++ {
				allRows[numStrings2+i] = append([]byte{'-'}, allRows[numStrings2+i]...)
			}
			row--
		} else {
			//DIAG
			row--
			col--
		}
	}

	// while either row or col is positive, backtrack to source

	for row > 0 {

	}

	for col > 0 {

	}

	// we have a big 2-D slice of bytes that we need to convert to a slice to strings
	for k := range alignment {
		//k-th row becomes string conversion of k-th row of byte sice
		alignment[k] = string(allRows[k])
	}

	return alignment
}
