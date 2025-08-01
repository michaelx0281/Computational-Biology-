package main

// GreedyAssembler takes a collection of strings and returns a genome whose
// k-mer composition is these strings. It makes the following assumptions.
// 1. "Perfect coverage" -- every k-mer is detected
// 2. No errors in reads
// 3. Every read has equal length (k)
// 4. DNA is single-stranded
func GreedyAssembler(reads []string) string {
	reads2 := reads
	genome := reads2[0]
	len_genome := len(genome)
	reads2 = reads[1:] //?? why does it remove the first str?
	for len(reads2) > 0 {
		for i := 0; i < len(reads2); i++ {
			if HasPrefix(reads2[i], genome, len_genome, len(reads2[i])) { //has suffix is not the same thing as equal to
				genome += string(reads2[i][len(reads2[i])-1])
				len_genome++
				reads2 = RemoveRead(reads2, reads2[i])

			}
		}

		for i := 0; i < len(reads2); i++ {
			if HasSuffix(reads2[i], genome, len(reads2[i])) {

				genome = string(reads2[i][0]) + genome
				len_genome++
				reads2 = RemoveRead(reads2, reads2[i])

			}
		}
	}

	return genome

}

func RemoveRead(reads []string, read string) []string {
	for i, r := range reads {
		if r == read {
			reads = append(reads[:i], reads[i+1:]...)
			break
		}
	}
	return reads
}

// I love this
func HasSuffix(read string, genome string, len int) bool {
	if read[1:] == genome[0:(len-1)] {
		return true
	}
	return false
}

func HasPrefix(read string, genome string, len_genome int, len_read int) bool {
	if read[:len_read-1] == genome[len_genome-len_read+1:] {
		return true
	}

	return false
}

//this method does not tend to work out well when a lot of the reads have high redudancy substring portions
//the non-diagonal reads may have scores of 0 or 1
//the main diagonal reads have giant values but are rendered to zero because they are not interesting / relevant to our problem
