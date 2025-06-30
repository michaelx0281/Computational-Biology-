package main

import "fmt"

// LCSLength takes two strings as input. It returns the length of a longest common subsequence of the two strings.
func LCSLength(str1, str2 string) int {

	matrix := LCSMatrix(str1, str2)
	return matrix[len(str1)][len(str2)]
}

type Pair struct {
	i int
	j int
}

//LCSPaths takes two strins as input. It returns the number of paths that could be taken to find the longest commom subsequence of the two strings
/*

This will use the concept of recursion with the addition of concepts from fasciculation--when new paths find the pioneering (old) paths they automatically know that a path (which follows the heaviest weight criteria) to the source is possible. Inversely, an opposite signal would be make using a variable that indicates a path definitely does not lead to the sink (and turn further signals away).

I am borrowing these notions from neuroscience signalling during development in terms of attractive and repulsive chemical signalling molecules

*/
func LCSPaths(str1, str2 string, i, j int) int {
	matrix := LCSMatrix(str1, str2)

	dMatrix := Make2D_2[bool](len(matrix), len(matrix[0]))
	fMatrix := Make2D_2[bool](len(matrix), len(matrix[0]))

	//BaseCase(s)
	//if at the source, return 1
	if i == 0 && j == 0 {
		return 1
	} else if fMatrix[i][j] == true {
		return 1
	} else if dMatrix[i][j] == true {
		return 0
	}

	// first step: find out if the piece diagonally up has a smaller numerical value
	if matrix[i-1][j-1] < matrix[i][j] {
		paths := LCSPaths(str1, str2, i-1, j-1)

		// if condition is fufilled, temporary matrix and recording matrix both filled
		if paths == 0 {
			dMatrix[i][j] = true
			DeadEnd(dMatrix, Pair{i: i, j: j}) //there may be some redundancy here
		}
		if paths > 0 {
			fMatrix[i][j] = true
			Fasciculate(fMatrix, Pair{i: i, j: j})
		}

		return paths
	} else if matrix[i-1][j] == matrix[i][j-1] {
		//tie!
		upPaths := LCSPaths(str1, str2, i-1, j)
		rightPaths := LCSPaths(str1, str2, i, j-1)

		if upPaths == 0 {
			dMatrix[i][j] = true
		} else if rightPaths == 0 {
			dMatrix[i][j] = true
		}

		if upPaths > 0 {
			fMatrix[i][j] = true
		} else if rightPaths > 0 {
			fMatrix[i][j] = true
		}

		return upPaths + rightPaths
	}

	// p := Pair{i: 1, j: 3}

	return 0
}

// func (p Pair) getI() int{
// 	return p.i
// }

// func (p Pair) getJ() int{
// 	return p.j
// }

//these two functions down here are able to keep count!

func Fasciculate(matrix [][]bool, p ...Pair) func() [][]bool {
	trueMatrix := Make2D_2[bool](len(matrix), len(matrix[0]))

	return func() [][]bool {
		for i := 0; i < len(p); i++ {
			trueMatrix[p[i].i][p[i].j] = true
		}

		return trueMatrix
	}
}

func DeadEnd(matrix [][]bool, p ...Pair) func() [][]bool {
	trueMatrix := Make2D_2[bool](len(matrix), len(matrix[0]))

	return func() [][]bool {
		for i := 0; i < len(p); i++ {
			// if trueMatrix[p[i].i][p[i].j] == true {
			// 	continue
			// }
			trueMatrix[p[i].i][p[i].j] = true
		}

		return trueMatrix
	}
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

func LCSMatrix(str1, str2 string) [][]int {
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

	return matrix
}
