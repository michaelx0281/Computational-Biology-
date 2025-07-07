package main

import "fmt"

/*
Extras!

Introduction to Multiple Alignment
Progressive Alignment
Scoring Multiple Alignments
Partial Order Alignment
A-Bruijn Approach to Multiple Alignment

Up until THIS POINT we have only tried to align TWO sequences

What about aligning more than two?

There is more significance once a pattern between two sequences appears in many others.

Multiple alignments can reveal subtle similarities that pairwise alignments do not reveal.

Alignment of 2 sequences is represented as a 2-row matrix. In a similar way, we represent alignment of 3 sequences as 3-row matrix

A T - G C G -
A - C G T - A
A T C A C _ A

Our scoring function should score alignments with conserved columns higher.

Three sequences to align: ATGC, AATC, ATGC

The coordinates should be plotted in a 3D space!

The same strategy could be applied as the one aligning two sequences.

A 3-D Manhattan Cube with each axis represeting a sequence to align

For global alignments --> source to sink

2D Alignment Cell versus a 3D Alignment Cell

There are 3 edges in each unit square (2D)
There are 7 edges in each unit cube (3D)

For the hyperdimensional planes of MSA matrix comparisons, there are a total of 2^n - 1 vertices to account for (traveling from the source to the sink)

Each different type of direction would be taken as a different input!!

For example, in the 2D representation, there are 4=1 different representations (aka 3) (right, down, diag)

For this specific 3D illustration, there is a grand total of 7!!!!
*/

func MultipleGlobleAlignmentScoreTable(str1, str2, str3 string) [][][]float64 {
	//somehow, there needs to be a defined:
	/*
		X-axis
		Y-axis
		Z-axis

		(xyz) == sink!!!

		(000) == source!!!

		Lets find / list the other 7 combinations!!

		(here it is in terms of i,j,k instead)

		i-1, j-1, k-1 (this is called the "cube diagonal") +2

		i-1, j-1, k (these are called the face diagonals) +1
		i-1, j, k-1
		i, j-1, k-1

		i-1, j, k (these are called the edge diagonals) -1
		i, j-1, k
		i, j, k-1

		for 3 sequences of length n, the run time is 7n^3

		seq v. seq

		how to align seq. v. profile??

		I have now seen the representation of it, and it is quite clever. I have also seen more of the model of how it should behave when aligning two different profiles together as well.

		You can use decimal or percentage weights that helps standardize the amount of match between each of the different sequences that you are currently analyzing.

		First of all, I would like to definitely want to create a program / function which is able to create a profile from two similar sequences. Lets take this a step at a time (you are smart!! don't doubt yourself!!)

	*/

	face_mtx := Make2D_2[float64](len(str1)+1, len(str2)+1)
	cube_mtx := Make3D[float64](len(str3)+1, face_mtx) //this is a matrix of type [][][]int, essentially

	c_diagMatch := 1 //I am not sure if these are the best values to use
	c_diagMisMatch := -1
	f_diagMatch := 1
	f_diagMisMatch := -2
	e_diag := -3

	for i := 0; i < len(str3); i++ {
		for j := 1; j < len(str2); j++ {
			for k := 1; k < len(str1); k++ {
				//in all of these cases, make sure that the indexes are actually within bounds

				//the below are all of the matches (which are varying degrees of diagonal)

				//this is the case where the cube diagonal matches, +1

				if i == 0 {
					continue
				}
				if (str1[i-1] == str2[j-1]) && (str2[j-1] == str3[k-1]) {
					cube_mtx[i][j][k] = cube_mtx[i-1][j-1][k-1] + float64(c_diagMatch)
				} else if str1[i-1] == str2[j-1] || str2[j-1] == str3[k-1] || str3[k-1] == str1[i-1] { // this is the case of having 1 indel
					cube_mtx[i][j][k] = cube_mtx[i-1][j-1][k-1] + float64(f_diagMatch)
				} else {
					cube_mtx[i][j][k] = cube_mtx[i-1][j-1][k-1] + float64(e_diag)
				}

				//this is the case for any mismatches of any sort
				cube_mtx[i][j][k] = max(
					float64(cube_mtx[i][j-1][k-1])+float64(f_diagMisMatch),
					float64(cube_mtx[i-1][j-1][k])+float64(f_diagMisMatch),
					float64(cube_mtx[i-1][j][k-1])+float64(f_diagMisMatch),
					float64(cube_mtx[i-1][j-1][k-1])+float64(c_diagMisMatch),
				)
			}
			fmt.Println(cube_mtx[i][j])
		}
		fmt.Println(cube_mtx[i])
	}

	return cube_mtx
}

type ThreeAlignment [3]string

/*
TripleAlignment should take in strings str1, str2, and str3 and return as output the weighted alignment of the strings
*/

func TripleAlignment(str1, str2, str3 string) ThreeAlignment {
	var alignment ThreeAlignment

	c_diagMatch := 1 //I am not sure if these are the best values to use
	c_diagMisMatch := -1
	f_diagMatch := 1
	f_diagMisMatch := -2
	e_diag := -3

	matrix := MultipleGlobleAlignmentScoreTable()
	score_matrix := Make3D[float64]( len(str3), Make2D_2[float64](len(str1), len(str2)))

	position := matrix[len(str1)][len(str2)][len(str3)]

	row := len(str1)
	col := len(str2)
	width := len(str3)

	array1 := str1[row]
	array2 := str2[col]
	array3 := str3[width]

	for row != 0 && col != 0 && width != 0 {
		var predecessor_value float64
		
		var store []float64

		if (str1[row-1] == str2[col-1]) && (str2[col-1] == str3[width-1]) { //the cases for matches? 
			store = append(store, float64(matrix[row-1][col-1][width-1]) + float64(c_diagMatch))
		} else if str1[row-1] == str2[col-1] || str2[col-1] == str3[width-1] || str3[width-1] == str1[row-1] { // this is the case of having 1 indel
			store = append(store, float64(matrix[row-1][col-1][width-1]) + float64(f_diagMatch))
		}

		mismatched_val:=  //if there are no matches, then the current point value would be different?
			max(
				matrix[row-1][col-1][width-1] + float64(c_diagMisMatch),
				matrix[row][col-1][width-1] + float64(f_diagMisMatch),
				matrix[row-1][col][width-1] + float64(f_diagMisMatch),
				matrix[row-1][col][width] + float64(f_diagMisMatch),
				matrix[row][col-1][width] + float64(e_diag),
				matrix[row][col][width-1] + float64(e_diag),
				matrix[row-1][col-1][width-1] + float64(e_diag),
			)
		 
		for _, value := range store {
			predecessor_value = max(value, mismatched_val)
		}

		if matrix[row-1][col-1][width-1] + float64(c_diagMatch) == predecessor_value {
			str1 = append(str1, str1[row-1])
			str2 = append(str2, str2[col-1])
			str3 = append(str3, str3[width-1])
			row--
			col--
			width--
		} else if matrix[row-1][col-1][width] + float64(f_diagMatch)
	}

	return alignment
}

type NSeqProfile []string

func (profile NSeqProfile) ConsesusSequence() string {
	consensusSeq := make([]byte, len(profile[0]))

	for i, seq := range profile {
		for j, char := range seq {
			if i == 0 {
				consensusSeq[j] = byte(char)
			} else if consensusSeq[j] != byte(char) {
				consensusSeq[j] = '~'
			}
		}
	}

	return string(consensusSeq)
}

func makeProfile(sequences ...string) NSeqProfile {
	var profile NSeqProfile

	for _, seq := range sequences {
		profile = append(profile, seq)
	}

	return profile
}

func AlignSeqToProfile(seq string, profile NSeqProfile) (string, float64) { //this is not the optimal way to do this, but this works sorta?
	/*
		I need to make a scoring system such that each of the kth-index sequences within the profile.

		Previously, the way that this worked was that there was a full point given for any of the matches. Now, there much more of
		a weighted system that provides 1/k-sequences points for each match of characters from the sequence and characters from each sequence
		inside the profile.

	*/

	match := float64(1) / float64(len(profile))
	mismatch := -1.0
	gap := -3.0

	score := 0.0

	for i := range len(profile) {
		score += PairGlobalAlignmentScore(seq, profile[i], match, mismatch, gap)
	}

	score /= float64(len(profile))

	profile = append(profile, seq)
	consensusSeq := profile.ConsesusSequence()

	return consensusSeq, score

}

//I need a function which runs something which is similar to GlobalAlignment
func PairGlobalAlignmentScore(str1, str2 string, match, mismatch, gap float64) float64 {
	mtx := GlobalScoreTable(str1, str2, match, mismatch, gap)
	return mtx[len(mtx)-1][len(mtx[0])-1]
}
