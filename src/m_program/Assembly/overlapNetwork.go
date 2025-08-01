package main

// MakeOverlapNetwork() takes a slice of reads with match, mismatch, gap and a threshold.
// It returns adjacency lists of reads; edges are only included
// in the overlap graph is the alignment score of the two reads is at least the threshold.
func MakeOverlapNetwork(reads []string, match, mismatch, gap, threshold float64) map[string][]string {
	//Initialize an adjacency list to represent the overlap graph.
	// adjacencyList := make(map[string][]string)

	//make the overlap matrix
	overlapMatrix := OverlapScoringMatrix(reads, match, mismatch, gap)

	//binarize it
	binarizedMatrix := BinarizeMatrix(overlapMatrix, threshold)

	//convert it to an adjacency list
	return ConvertAdjancencyMatrixToList(reads, binarizedMatrix)
}

// ConvertAdjancencyMatrixToListj
// Input: a binary matrix representing an adjacency matrix
// Output: a list representing the adjacency list
func ConvertAdjancencyMatrixToList(reads []string, b [][]int) map[string][]string {
	adjacencyList := make(map[string]([]string))

	// range over the entire matrix, and any time you see a 1, add the appropriate edte to the adjacnecy list
	for r := range b {
		for c := range b[r] {
			if b[r][c] == 1 {
				// we are currently entering the reads that are connected from reads[r]
				// we found the reads[r] should be connected to reads[c]
				currentRead := reads[r]
				adjacencyList[currentRead] = append(adjacencyList[currentRead], reads[c])

			}
		}
	}

	return adjacencyList
}
