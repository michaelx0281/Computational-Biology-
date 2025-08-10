package main

//I think that this Charging Station promopt is possibly one of the hardest / harder problems that I had to do!

//import "github.com/michaelx0281/Computational-Biology/src/utils"

/*

Neighbors(Pattern, d)
    if d = 0
        return {Pattern}
    if |Pattern| = 1
        return {A, C, G, T}
    Neighborhood ← an empty set
    SuffixNeighbors ← Neighbors(Suffix(Pattern), d)
    for each string Text from SuffixNeighbors
        if HammingDistance(Suffix(Pattern), Text) < d
            for each nucleotide x
                add x • Text to Neighborhood
        else
            add FirstSymbol(Pattern) • Text to Neighborhood
    return Neighborhood

*/

func Neighbors(Pattern string, d int) []string {
	neighborhood := make([]string, 0)

	if d == 0 {
		neighborhood = append(neighborhood, Pattern)
		return neighborhood
	}

	if len(Pattern) == 1 {
		neighborhood = append(neighborhood, []string{"A", "C", "G", "T"}...) //it seems like i've got this notation down now!
		return neighborhood
	}

	SuffixNeighbors := Neighbors(Suffix(Pattern), d)

	nucleotides := Nucleotides()

	for _, text := range SuffixNeighbors {
		if HammingDist(Suffix(Pattern), text) < d {
			for _, nucleotide := range nucleotides {
				neighborhood = append(neighborhood, nucleotide+text) //idk if this is the prettiest way to do it, but if it works, it works. --> this adds a,g,c,t with each possible variation of suffixes at each length
			}
		} else {
			neighborhood = append(neighborhood, string(Pattern[0])+text)
		}
	}

	return neighborhood
}

func Suffix(Pattern string) string {
	return Pattern[1:]
}

func Nucleotides() []string {
	return []string{"A", "T", "C", "G"}
}
