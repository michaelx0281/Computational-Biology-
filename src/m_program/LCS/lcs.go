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

//double check the way that this works and make sure that everything here seems to work correctly
func LCSPaths(str1, str2 string, i, j int) int {
	matrix := LCSMatrix(str1, str2)

	fakeDMatrix := matrix
	fakeFMatrix := matrix

	dMatrix := Make2D_2[bool](len(matrix), len(matrix[0]))
	fMatrix := Make2D_2[bool](len(matrix), len(matrix[0]))

	dMatrix = DeadEnd(dMatrix, Pair{i: 0, j: 0})()
	for i := range dMatrix {
		for j := range dMatrix[0] {
			if dMatrix[i][j] == true {
				fakeDMatrix[i][j] = -1
			}
		}
	}
	fMatrix = Fasciculate(fMatrix, Pair{i: 0, j: 0})()

	for i := range fMatrix {
		for j := range fMatrix[0] {
			if fMatrix[i][j] == true {
				fakeFMatrix[i][j] = -1
			}
		}
	}

	//BaseCase(s)
	//if at the source, return 1
	if i == 0 && j == 0 {
		return 1
	} else if fMatrix[i][j] == true {
		return 1
	} else if dMatrix[i][j] == true {
		return 0
	}

	if i == 0 || j == 0 {
		return 1
	}

	// first step: find out if the piece diagonally up has a smaller numerical value
	if matrix[i-1][j-1] < matrix[i][j] {
		paths := LCSPaths(str1, str2, i-1, j-1)
		// if condition is fufilled, temporary matrix and recording matrix both filled
		if paths == 0 {
			dMatrix[i][j] = true
			dMatrix = DeadEnd(dMatrix, Pair{i: i, j: j})() //there may be some redundancy here
		}
		if paths > 0 {
			fMatrix[i][j] = true
			fMatrix = Fasciculate(fMatrix, Pair{i: i, j: j})()
		}

		return paths
	} else if matrix[i-1][j] == matrix[i][j-1] {
		//tie!
		upPaths := LCSPaths(str1, str2, i-1, j)
		rightPaths := LCSPaths(str1, str2, i, j-1)

		if upPaths == 0 {
			dMatrix[i][j] = true
			dMatrix = DeadEnd(dMatrix, Pair{i: i, j: j})()
		} else if rightPaths == 0 {
			dMatrix[i][j] = true
			dMatrix = DeadEnd(dMatrix, Pair{i: i, j: j})()
		}

		if upPaths > 0 {
			fMatrix[i][j] = true
			fMatrix = Fasciculate(fMatrix, Pair{i: i, j: j})()
		} else if rightPaths > 0 {
			fMatrix[i][j] = true
			fMatrix = Fasciculate(fMatrix, Pair{i: i, j: j})()
		}

		return upPaths + rightPaths
	}

	// p := Pair{i: 1, j: 3}

	return 0
}

//backtracking pointers
//finding LCS in general

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
// Output: longest (int terms of weight) paths

func Make2D_2[T any](n, m int) [][]T {
	matrix := make([][]T, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]T, m)
	}

	return matrix
}

func Make3D[T any](n int, m [][]float64) [][][]T {
	matrix := make([][][]T, n)

	for i := 0; i < n; i++ {
		matrix[i] = Make2D_2[T](len(m), len(m[0]))
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
			matrix[i][j] = max(matrix[i][j-1], matrix[i-1][j], matrix[i-1][j-1])

		}
		fmt.Println(matrix[i])
	}

	return matrix
}

/*
The edit distance is the minimum number of mismatches, deletions, and insertions that are needed.
This measures how good the alignment is. What is the fewest number of changes that it takes to go from one string to the other?

Minimizing the number of mismatches is a completely different problem!
*/

/*
Reconstructing the Optimal Path!!

Every single time that you find an match edge, you can add one to the count! otherwise, if there is not match edge, do not add points
*/

func LCS(str1, str2 string) int {

	return 0
}

func EditDistance(str1, str2 string) int {

	matrix := InitializeMatrix(Make2D_2[int](len(str1)+1, len(str2)+1)) //str1 is the col indicies, and str2 is the row indicies

	// for each cell in the matrix, check first if the two letters are equal to eachother, then add one to
	// the value of the cell, if they are not the same, then take the bigger of the two adjacent cells

	//i is the col
	//j is the row

	for i := 1; i <= len(str1); i++ {
		for j := 1; j <= len(str2); j++ {
			if str1[i-1] == str2[j-1] {
				//this checked of there is an alignment match
				matrix[i][j] = matrix[i-1][j-1]
			} else {
				//this is the second
				matrix[i][j] = min(matrix[i-1][j-1]+1, min(matrix[i][j-1]+1, matrix[i-1][j]+1))
			}
			//fmt.Println(matrix)
		}
	}

	return matrix[len(str1)][len(str2)]
}

// Input: indices i and j which you want to find --> input the sink
// Output: longest (int terms of weight) path

func InitializeMatrix(mtx [][]int) [][]int {
	for row := 1; row < len(mtx); row++ {
		mtx[row][0] = row
	}

	for col := 1; col < len(mtx[0]); col++ {
		mtx[0][col] = col
	}
	return mtx
}

func EditDistanceMatrix(patterns []string) [][]int {

	matrix := Make2D_2[int](len(patterns), len(patterns)) //str1 is the col indicies, and str2 is the row indicies

	// the value of the cell, if they are not the same, then take the bigger of the two adjacent cells

	//i is the col
	//j is the row

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if j != i && matrix[i][j] == 0 {
				matrix[i][j] = EditDistance(patterns[i], patterns[j])
				matrix[j][i] = EditDistance(patterns[i], patterns[j])
			}
		}
	}

	return matrix
}

func InitializeMatrix2(mtx [][]float64, gap float64) [][]float64 {
	for row := 1; row < len(mtx); row++ {
		mtx[row][0] = float64(row) * gap
	}

	for col := 1; col < len(mtx[0]); col++ {
		mtx[0][col] = float64(col) * gap
	}
	return mtx
}

func GlobalScoreTable(str1, str2 string, match, mismatch, gap float64) [][]float64 {

	matrix := InitializeMatrix2(Make2D_2[float64](len(str1)+1, len(str2)+1), gap*-1) //str1 is the col indicies, and str2 is the row indicies

	//fmt.Println("LCS Matrix", matrix)

	// for each cell in the matrix, check first if the two letters are equal to eachother, then add one to
	// the value of the cell, if they are not the same, then take the bigger of the two adjacent cells

	//i is the col
	//j is the row

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
		//fmt.Println(matrix[i])
	}

	return matrix
}

func LongestCommonSubsequence(str1, str2 string) string {

	matrix := LCSMatrix(str1, str2)
	lcs := ""
	/**
		for i:=len(str1); i>=0; i--{
			for j:=len(str2); j>=0; j--{
				if (matrix[i][j]-1==matrix[i-1][j-1]){
					lcs=string(str1[i-1])+lcs
				} else if (matrix[i][j]==matrix[i][j-1]){
	                lcs=string(str1[i-1])+lcs
	            } else if (matrix[i][j]==matrix[i-1][j]){
	                lcs=string(str2[j-1])+lcs
	            }
			}
		}
	    **/
	col := len(str2)
	row := len(str1)
	for true {
		if row <= 0 && col <= 0 {
			break
		}
		if (row-1) >= 0 && (col-1) >= 0 && matrix[row][col]-1 == matrix[row-1][col-1] && str1[row-1] == str2[col-1] {
			lcs = string(str1[row-1]) + lcs
			col--
			row--

		} else if ((col - 1) >= 0) && matrix[row][col] == matrix[row][col-1] {
			col--
		} else if ((row - 1) >= 0) && matrix[row][col] == matrix[row-1][col] {
			row--
		}

	}
	return lcs
}

/*
Would you rather align two genes as DNA strings (nucleotides) or as proteins (amino acids)?

They saw the writing on the wall -- they would like to have the algorithm ready for when they were ready to sequence the genome

Would you rather align 2 DNA strings (nucleotide pairs) or align as proteins (amino acids)?

--> proteins seem like they would be a safer bet for some purposes, however, I am also interested in the regulatory roles of different regions of DNA like within the UTR, promoters, enhancers, and intron regions.

--> if you only care about what the protein does, that could be better for structural and functional comparisions (redundancy allows for more leeway)

--> sequencing DNA and proteins are very different
		--> setting up the information (even if it changes at dna level and not protein level, still interesting to look at)
--> align two things with 4 letters or 20 letters --> 20 tends to provide more information
--> protein strings at least 3 times shorter

only one or two percent encodes proteins

if you know the gene goes to the protein level, the proteins are more informative


*/

type Alignment [2]string

/*
Differences in Hemoglobin

- Common ancestry
- Different levels of oxygen and (aq) --> O2 (aq) might influence favored traits + pressure (selective pressures)
- Hemoglobin dual function?
- N - binding? O2 not separated from one another (differences in circulation)


For the case of the sars spike protein, there is a lot more limit in terms of space. This necesitates a change in the weights :)



*/

// Insert your GlobalAlignment() function here, along with any subroutines that you need. Please note the subroutines indicated in the problem description that are provided for you.
func GlobalAlignment(str1, str2 string, match, mismatch, gap float64) Alignment {
	var alignment Alignment
	mtx := GlobalScoreTable(str1, str2, match, mismatch, gap)

	// create the first indicies based on the length of the two strings
	row := len(str1)
	col := len(str2)

	// loop until you hit the left or top edge
	for row != 0 || col != 0 {
		// find the values up, left, and diagonal from the current index
		left := -1000.0
		if InMatrix(mtx, row, col-1) {
			left = mtx[row][col-1] - gap
		}

		up := -1000.0
		if InMatrix(mtx, row-1, col) {
			up = mtx[row-1][col] - gap
		}

		diag := -1000.0
		if InMatrix(mtx, row-1, col-1) {
			diag = mtx[row-1][col-1]
			if str1[row-1] == str2[col-1] {
				diag += match
			} else {
				diag -= mismatch
			}
		}

		// if there's a match use it! Otherwise go up or left if they are the same
		if diag == mtx[row][col] {
			alignment[0] = string(str1[row-1]) + alignment[0]
			alignment[1] = string(str2[col-1]) + alignment[1]
			row--
			col--
		} else if up == mtx[row][col] {
			alignment[0] = string(str1[row-1]) + alignment[0]
			alignment[1] = "-" + alignment[1]
			row--
		} else if left == mtx[row][col] {
			alignment[0] = "-" + alignment[0]
			alignment[1] = string(str2[col-1]) + alignment[1]
			col--
		}

	}
	return alignment
}

func InMatrix(matrix [][]float64, row int, col int) bool {
	if row < 0 || col < 0 {
		return false
	}
	if row >= len(matrix) {
		return false
	}
	if col >= len(matrix[row]) {
		return false
	}
	return true
}

func LocalScoreTable(str1, str2 string, match, mismatch, gap float64) [][]float64 {

	matrix := Make2D_2[float64](len(str1)+1, len(str2)+1) //str1 is the col indicies, and str2 is the row indicies

	//fmt.Println("LCS Matrix", matrix)

	// for each cell in the matrix, check first if the two letters are equal to eachother, then add one to
	// the value of the cell, if they are not the same, then take the bigger of the two adjacent cells

	//i is the col
	//j is the row

	for i := 1; i < len(str1)+1; i++ {
		for j := 1; j < len(str2)+1; j++ {
			if str1[i-1] == str2[j-1] {
				//this checked of there is an alignment match
				matrix[i][j] = matrix[i-1][j-1] + match
				continue
			}
			//this is the second case which takes the largest of the two previous
			matrix[i][j] = max(matrix[i][j-1]-gap, matrix[i-1][j]-gap, matrix[i-1][j-1]-mismatch)

			if matrix[i][j] < 0 {
				matrix[i][j] = 0
			}

		}
		//fmt.Println(matrix[i])
	}

	return matrix
}

func LocalAlignment(str1, str2 string, match, mismatch, gap float64) (Alignment, int, int, int, int) {
	var alignment Alignment
	mtx := LocalScoreTable(str1, str2, match, mismatch, gap)

	// create the first indicies based on the length of the two strings
	max_val := 0.0
	row := 0
	col := 0

	for i := 1; i <= len(str1); i++ {
		for j := 1; j <= len(str2); j++ {
			if mtx[i][j] > max_val {
				row = i
				col = j
				max_val = mtx[i][j]
			}
		}
	}

	endStr1 := row
	endStr2 := col

	// loop until you hit the left or top edge
	for mtx[row][col] > 0 {
		// find the values up, left, and diagonal from the current index
		left := -1000.0
		if InMatrix(mtx, row, col-1) {
			left = mtx[row][col-1] - gap
		}

		up := -1000.0
		if InMatrix(mtx, row-1, col) {
			up = mtx[row-1][col] - gap
		}

		diag := -1000.0
		if InMatrix(mtx, row-1, col-1) {
			diag = mtx[row-1][col-1]
			if str1[row-1] == str2[col-1] {
				diag += match
			} else {
				diag -= mismatch
			}
		}

		// if there's a match use it! Otherwise go up or left if they are the same
		if diag == mtx[row][col] {
			alignment[0] = string(str1[row-1]) + alignment[0]
			alignment[1] = string(str2[col-1]) + alignment[1]
			row--
			col--
		} else if up == mtx[row][col] {
			alignment[0] = string(str1[row-1]) + alignment[0]
			alignment[1] = "-" + alignment[1]
			row--
		} else if left == mtx[row][col] {
			alignment[0] = "-" + alignment[0]
			alignment[1] = string(str2[col-1]) + alignment[1]
			col--
		}

	}
	return alignment, row, endStr1, col, endStr2
}

//look into fitting alignment and overlapping alignment

//monte carlo simulation to find best parameters for given situation?
/*
What does it mean for a parameter to be good? I would get the sequence that is available with the least amount of gaps possible
--> have a separate overall ranking for this!

change the parameters to find the best results!

in practice, there are 20-30+ parameters

start with data that you know, how often does it find that specific region? --> classification in a sense

Scoring matrix; penalizes different matches / mismatches differently

BLOSUM62 Matrix --> very common and used in bioinformatics ``
*/

/*

Why would you reward or penalize different pairs differently?

--> classification usages --> if you are looking for specific concentration of a property of the Proteins / paremters aligned

How did they come up with this?

--> people were looking to find a specific property of these proteins

--> have a ranking based on how well a specific protein would fit that property (clustering into groups)

--> select scores + run

How would this change our alignment?

--> this is more focused and specific towards the selected properties / selective pressures

--> its like a mini-evolution

Can also sort for motifs and for bigger things too like folds and specific groups (steric and otherwise)

My question is: Why does the BLOSUM62 matrix look like a simplex / triangle instead of a manhattan plot?

Specific protein folding?

Determinance of protein function due to mutation changes / others
--> how similar proteins are

How replacable a protein is? Weights? Each are parameters.

if some mismatches are more common based on this matrix, don't penalize as much.const
if they are less common, penalize them more.

you can change the edge weights in order to change the network


Affine penalty: a way of socring contiguous gaps higher than discontiguous gaps.const
gap opening penalty --> less contiguous
gap extension penalty --> more contiguous

Alignment with affine gap penalties problem


Add a huge number of new gaps that are longer?

Divide the network into three layers
-> gap => jump out of the middle layer

3-layer manhattan plot

if you have so many different layers, could you try to run different alignments on each layer?
Do the three layer ones.const

Multiple Alignment Problem
	

*/
