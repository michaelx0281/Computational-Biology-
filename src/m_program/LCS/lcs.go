package main

import "fmt"

// LCSLength takes two strings as input. It returns the length of a longest common
// subsequence of the two strings.
func LCSLength(str1, str2 string) int {

	matrix := Make2D_2[int](len(str1)+1, len(str2)+1) //str1 is the col indicies, and str2 is the row indicies

	fmt.Println("LCS Matrix", matrix)

	// for each cell in the matrix, check first if the two letters are equal to eachother, then add one to
	// the value of the cell, if they are not the same, then take the bigger of the two adjacent cells

	//i is the col
	//j is the row

	for i := 0; i < len(str1)+1; i++ {
		for j := 1; j < len(str2)+1; j++ {
			if i == 0 {
				matrix[0][j] = 0
				continue
			}
			if str1[i-1] == str2[j-1] {
				//this checked of there is an alignment match
				matrix[i][j] = matrix[i-1][j-1] + 1
				continue
			}
			//this is the second case which takes the largest of the two previous
			matrix[i][j] = max(matrix[i][j-1], matrix[i-1][j])

		}
		fmt.Println(matrix[i])
	}

	return matrix[len(str1)][len(str2)]
}

// Input: indices i and j which you want to find --> input the sink
// Output: longest (int terms of weight) path

func Make2D_2[T any](n, m int) [][]T {
	matrix := make([][]T, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]T, m)
	}

	return matrix
}
