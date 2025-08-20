package main

//note that nothing here in this document is tested yet, this will just serve as more of a intro / personal project for now!

//There are a bunch of other things that I really want to put more focus on at the current moment so this stuff would probably have to wait until later!

import "math"

type col [4]float64

func CalculateEntropy(motifMtx []col) float64 {
	sum := 0.0
	//take it col by col
	for _, col := range motifMtx {
		temp := 0.0
		for _, entry := range col {
			temp += entry * math.Log2(entry)
		}

		//adding entropy of colum into collective sum
		sum += -1 * temp
	}

	return sum
}

func ProfileMtxToColMtx(p []Profile, length int) []col {
	mtx := make([]col, length)
	for i, profile := range p {
		mtx[i][0] = profile.A
		mtx[i][1] = profile.T
		mtx[i][2] = profile.C
		mtx[i][3] = profile.G
	}

	return mtx
}

type colDistribution struct {
	A int
	T int
	C int
	G int
}

func DistributionMatrix(Dna []string) ([]colDistribution, int) {
	//assume strings of equal length within the DNA ofc
	distributionMatrix := make([]colDistribution, len(Dna[0]))

	for _, strand := range Dna {
		for j, nuc := range strand {
			n := string(nuc)
			switch n {
			case "A":
				distributionMatrix[j].A++

			case "T":
				distributionMatrix[j].T++

			case "C":
				distributionMatrix[j].C++

			case "G":
				distributionMatrix[j].G++
			}
		}
	}

	return distributionMatrix, len(Dna[0])
}

type Profile struct {
	A float64
	T float64
	C float64
	G float64
}

func (p *Profile) DistToProfile(dist colDistribution) {
	total := totalN(dist)

	p.A = float64(dist.A) / float64(total)
	p.T = float64(dist.T) / float64(total)
	p.C = float64(dist.C) / float64(total)
	p.G = float64(dist.G) / float64(total)
}

func totalN(dist colDistribution) int {
	return dist.A + dist.C + dist.T + dist.G
}

func ProfileMatrix(dist []colDistribution, length int) ([]Profile, int) {
	profile := make([]Profile, length)
	for i, col := range dist {
		profile[i].DistToProfile(col)
	}

	return profile, length
}
