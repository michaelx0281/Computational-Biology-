package main

// "fmt"
// "math"

//ALL PENALTIES POSITIVE

// ScoreOverlapAlignment takes two strings along with match, mismatch, and gap penalties.
// It returns the maximum score of an overlap alignment in which str1 is overlapped with str2.
// Assume we are overlapping a suffix of str1 with a prefix of str2.
func ScoreOverlapAlignment(str1, str2 string, match, mismatch, gap float64) float64 {

	matrix := Make2D_2[float64](len(str1)+1, len(str2)+1)
	for i := 0; i <= len(str1); i++ {
		matrix[i][0] = 0
	}
	for i := 1; i < len(str1)+1; i++ {
		for j := 1; j < len(str2)+1; j++ {
			if str1[i-1] == str2[j-1] {
				//this checked of there is an alignment match
				matrix[i][j] = matrix[i-1][j-1] + match
				continue
			}
			//this is the second case which takes the largest of the two previous
			matrix[i][j] = max(matrix[i][j-1]-gap, matrix[i-1][j]-gap, matrix[i-1][j-1]-mismatch)
		}
	}

	maxScore := 0.0
	for j := 0; j <= len(str2); j++ {
		if matrix[len(str1)][j] > maxScore {
			maxScore = matrix[len(str1)][j]
		}
	}

	return maxScore

}

func Make2D_2[T any](n, m int) [][]T {
	matrix := make([][]T, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]T, m)
	}

	return matrix
}
