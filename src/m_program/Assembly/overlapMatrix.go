package main

import "fmt"

// OverlapScoringMatrix takes a collection of reads along with alignment penalties.
// It returns a matrix in which mtx[i][j] is the overlap alignment score of
// reads[i] with reads[j].
func OverlapScoringMatrix(reads []string, match, mismatch, gap float64) [][]float64 {
	mtx := Make2D_2[float64](len(reads), len(reads))
	for i := 0; i < len(reads); i++ {
		if i%10 == 0 {
			fmt.Println("Currently making row", i, "of overlap matrix")
		}
		for j := 0; j < len(reads); j++ {
			if i == j {
				mtx[i][j] = 0
				continue
			}
			mtx[i][j] = ScoreOverlapAlignment(reads[i], reads[j], match, mismatch, gap)
		}
	}
	return mtx

}
