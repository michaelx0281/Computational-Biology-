package main

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

func LCS3SA(str1, str2, str3 string) string {
	//somehow, there needs to be a defined:
	/*
		X-axis
		Y-axis
		Z-axis

		(xyz) == sink!!!

		(000) == source!!!

		Lets find / list the other 7 combinations!!

		(here it is in terms of i,j,k instead)

		i-1, j-1, k-1 (this is called the "cube diagonal")

		i-1, j-1, k (these are called the face diagonals)
		i-1, j, k-1
		i, j-1, k-1

		i-1, j, k (these are called the edge diagonals)
		i, j-1, k
		i, j, k-1

		for 3 sequences of length n, the run time is 7n^3

		seq v. seq

		how to align seq. v. profile??

		I have now seen the representation of it, and it is quite clever. I have also seen more of the model of how it should behave when aligning two different profiles together as well.

		You can use decimal or percentage weights that helps standardize the amount of match between each of the different sequences that you are currently analyzing.

		First of all, I would like to definitely want to create a program / function which is able to create a profile from two similar sequences. Lets take this a step at a time (you are smart!! don't doubt yourself!!)

	*/

	return ""
}

type NSeqProfile []string

func (profiles NSeqProfile) ConsesusSequence() string {
	consensusSeq := make([]byte, len(profiles[0]))

	for i, profile := range profiles {
		for j, char := range profile {
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
