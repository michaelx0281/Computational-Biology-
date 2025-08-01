package main

import "fmt"

func MakeOverlapNetworkMinimizers(reads []string, minimizerDictionary StringIndex, match, mismatch, gap, threshold float64) map[string][]string {

	numReads := len(reads)

	//this just makes the overlap matrix
	overlapMatrix := InitializeFloatMatrix(numReads, numReads)

	// everything is zero by default, which is kinda good
	bigNegative := threshold - 10000000.0

	//this flag will tell us have I overlapped two reads before?
	//at his point in time, that answer is NO for all pairs of reads

	for r := range overlapMatrix {
		for c := range overlapMatrix[r] {
			overlapMatrix[r][c] = bigNegative
		}
	}

	counter := 0

	// engine of function: range over minimizer map, and ask:
	// (1) do two strings share the same minmizer?
	// (2) have I overlapped them already?
	for _, readIndices := range minimizerDictionary {

		if counter%100 == 0 {
			fmt.Println("Now considering element", counter, "of minimizer map")
		}
		//readIndices is a slice of integers
		//range over this slice, and grab all possible pairs of elements that are in the slice (n choose 2)

		//how do you choose 2 out of each of the groups?
		//two moving frames. One moves every 'cycle' when the second observe gets to the end.

		for i := range readIndices {
			// fmt.Println("i: ", i)
			for j := 1; j < len(readIndices); j++ {
				// fmt.Println("Dict: ", minimizerDictionary)
				// note: mtx[i][j] != mtx[j][i]
				// fmt.Println("j: ", j)
				// fmt.Println("Length of readIndices", len(readIndices))

				index1 := readIndices[i]
				index2 := readIndices[j]

				read1 := reads[index1]
				read2 := reads[index2]

				//perform alignment of read1 and read2

				if overlapMatrix[index1][index2] == bigNegative {
					overlapMatrix[index1][index2] = ScoreOverlapAlignment(read1, read2, match, mismatch, gap)
				}

				if overlapMatrix[index2][index1] == bigNegative {
					overlapMatrix[index2][index1] = ScoreOverlapAlignment(read2, read1, match, mismatch, gap)
				}
			}
		}

		// fmt.Println("Updating counter.")
		//update counter
		counter++
	}

	//overlap matrix is now made. Binarize it!
	b := BinarizeMatrix(overlapMatrix, threshold)

	return ConvertAdjancencyMatrixToList(reads, b)
}

//use the transitive property

//if x -> y, y -> z, x -> z (get rid of this one)
