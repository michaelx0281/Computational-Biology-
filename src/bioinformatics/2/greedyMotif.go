package main

/*
	This isn't looking too great at the moment. Skip it if you can and come back to it later. There are more important things in the world :(
*/

// import "fmt"

// import "fmt"

/*
	This is supposed to exemplify the Greedy Motif Search.

	Profile-most Probable k-mer Problem: Find a Profile-most probable k-mer in a string.

	Input: A string Text, an integer k, and a 4 × k matrix Profile.
	Output: A Profile-most probable k-mer in Text.
*/

func ProfileMostProbableKmer(Text string, k int, Profile [4][]float64) string {
	//k represents the length of the kmer taken from text

	n := len(Text)

	maxProbability := 0.0
	motif := ""

	for i := 0; i < n-k+1; i++ {
		//now for each of these possibilities, you want to maximize the score of multiplying all of these different profiles together

		//ACGT in terms of order in the profile
		pattern := Text[i : i+k]

		distribution := make([]float64, k)

		for j, symbol := range pattern {
			s := byte(symbol)
			switch s {
			case 'A':
				if Profile[0][j] == 0 {
					distribution[j] = 0.01
					break
				} else if Profile[0][j] == 1 {
					distribution[j] = 0.99
					break
				}
				distribution[j] = Profile[0][j]
			case 'C':
				if Profile[1][j] == 0 {
					distribution[j] = 0.01
					break
				} else if Profile[1][j] == 1 {
					distribution[j] = 0.99
					break
				}
				distribution[j] = Profile[1][j]
			case 'G':
				if Profile[2][j] == 0 {
					distribution[j] = 0.01
					break
				} else if Profile[2][j] == 1 {
					distribution[j] = 0.99
					break
				}
				distribution[j] = Profile[2][j]
			case 'T':
				if Profile[3][j] == 0 {
					distribution[j] = 0.01
					break
				} else if Profile[3][j] == 1 {
					distribution[j] = 0.99
					break
				}
				distribution[j] = Profile[3][j]
			}
		}

		product := calcProbability(distribution)

		// fmt.Println("Product", i, product)

		if product > maxProbability {
			maxProbability = product
			// fmt.Println(maxProbability)
			motif = pattern

			// fmt.Println("Motif:", motif, "%", maxProbability)
		}
	}

	return motif
}

func calcProbability(distribution []float64) float64 {
	product := 1.0

	for _, float := range distribution {
		product *= float
	}

	return product
}

/*
Code Challenge: Implement GreedyMotifSearch().

Input: Integers k and t, followed by a space-separated collection of strings Dna.
Output: A collection of strings BestMotifs resulting from applying GreedyMotifSearch(Dna, k, t). If at any step you find more than one Profile-most probable k-mer in a given string, use the one occurring first.

Pseudocode:

GreedyMotifSearch(Dna, k, t)
    BestMotifs ← motif matrix formed by first k-mers in each string from Dna
    for each k-mer Motif in the first string from Dna
        Motif1 ← Motif
        for i = 2 to t
            form Profile from motifs Motif1, …, Motifi - 1
            Motifi ← Profile-most probable k-mer in the i-th string in Dna
        Motifs ← (Motif1, …, Motift)
        if Score(Motifs) < Score(BestMotifs)
            BestMotifs ← Motifs
    return BestMotifs
*/
//t = number of strings in Dna //assume standard to be ACGT
func GreedyMotifSearch(Dna []string, k, t int) []string {
	BestMotifs := GenerateVertMotifs(Dna, k, t)

	//horizontal motifs of first string
	n := len(Dna[0])
	for i := 0; i < n-k+1; i++ {
		//set to zero, because only want to check first strand only against the others
		motif := Dna[0][i : i+k]

		Motif := make([]string, 1)
		Motif[0] = motif

		for j := 1; j <= t-1; j++ { //there may be an error over here? i am not sure about why the bounds are why they are, but this may need some fixing in the future
			// fmt.Println(Motif)
			// fmt.Println(Dna[j])
			profile := ProfileMtx(Motif) //this is really weird and idk if this is the greatest idea
			Motif = append(Motif, ProfileMostProbableKmer(Dna[j], k, profile))

			// if j == t-1 {
			// 	fmt.Println(Motif)
			// }
		}

		if ScoreMtx(Motif) < ScoreMtx(BestMotifs) {
			BestMotifs = Motif
		}
	}

	return BestMotifs
}

// func ScoreMotif(motif []string) float64 {
// 	profile := ProfileMtx(motif)

// 	score := 1.0

// 	for i := 0; i < len(profile[0]); i++ {
// 		max := 0.0

// 		for j := 0; j < 4; j++ {
// 			if profile[j][i] > max {
// 				max = profile[j][i]
// 			}
// 		}

// 		score *= max
// 	}

// 	return score
// }

// //all of the work belew seems to be incorrect, now creating new functions

// //entropy is still experimental (aka i don't really want to test or touch it rn) --> lets make a smaller but simpler function creating profile

//should test next
func ScoreMtx(matrix []string) int { //it seems like to me, that I have been doing this wrong.. Need to revise to better fit the problem. Take the largest value out of each col, and multiply them all together
	n := len(matrix[0])
	// t := len(matrix)

	score := 0

	profile := ProfileMtx(matrix)

	dist := ProfileToDist(profile)

	for i := 0; i < n; i++ {
		score += ScoreCol([]int{dist[0][i], dist[1][i], dist[2][i], dist[3][i]})
	}

	return score
}

func ScoreCol(col []int) int {
	//find the maxInt amoungest these. Return the total minus the max.

	max := 0

	for i := 0; i < 4; i++ {
		if col[i] > max {
			max = col[i]
		}
	}

	//by this point we have found the max

	total := col[0] + col[1] + col[2] + col[3]

	return total - max // this is the score at the end
}

func ProfileToDist(profile [4][]float64) [4][]int {
	n := len(profile[0])
	t := len(profile)

	dist := make([][]int, 4)

	for i := range dist {
		dist[i] = make([]int, n)
	}

	for i := 0; i < 4; i++ {
		for j := 0; j < n; j++ {
			profile[i][j] *= float64(t)
			dist[i][j] = int(profile[i][j])
		}
	}

	return [4][]int(dist) //now updated --> this seems like a very dumb approach but I really do not want to start from scratch anymore, this file is also just a giant messs
}

//TESTED
func ProfileMtx(Dna []string) [4][]float64 { //i should test this out individually to see how well it works
	n := len(Dna[0])
	t := len(Dna)

	profile := make([][]float64, 4)

	for i := range profile {
		profile[i] = make([]float64, n)
	}

	//iterate horizontally
	for col := 0; col < n; col++ {
		//generate a list of all nucleotides in current col
		list := make([]byte, t)
		for row := 0; row < t; row++ {
			list[row] = byte(Dna[row][col])
		}

		for j := 0; j < 4; j++ {
			profile[j][col] = colCount(list)[j]
		}
	}

	return [4][]float64(profile)
}

func colCount(col []byte) [4]float64 { //hopefully this would work now?
	t := float64(len(col)) //this the the amount of rows to cover in the distribution
	list := CountNucleotides(string(col))

	newList := make([]float64, 4)

	for i := range list {
		newList[i] = float64(list[i]) / t
	}

	return [4]float64(newList)
}

//returns stuff in ACGT formatting
func CountNucleotides(Text string) []int {
	a := 0
	c := 0
	g := 0
	t := 0
	for i := 0; i < len(Text); i++ {
		n := byte(Text[i])

		switch n {
		case 'A':
			a++
		case 'C':
			c++
		case 'G':
			g++
		case 'T':
			g++
		}
	}

	return []int{a, c, g, t}
}

func GenerateVertMotifs(Dna []string, k, t int) []string {
	list := make([]string, t)

	for i := 0; i < t; i++ {
		list[i] = Dna[i][0:k]
	}

	return list
}
