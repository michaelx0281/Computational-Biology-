package main

import "github.com/michaelx0281/Computational-Biology/src/utils"

/*

Now, this is an oldie that I have done before! I have done did is around Chapter 1 or 2 of Programming for Lovers, let's revisit this and see how much of this I can remember!

I think that the early portions of the course certainly got a lot easier and streamlined for me!

Here is the pseudocode that comes with this section!

	FindClumps(Text, k, L, t)
    Patterns ← an array of strings of length 0
    n ← |Text|
    for every integer i between 0 and n − L
        Window ← Text(i, L)
        freqMap ← FrequencyTable(Window, k)
        for every key s in freqMap
            if freqMap[s] ≥ t
                append s to Patterns
    remove duplicates from Patterns
    return Patterns


	k the k-mer length that you are checking for
	L is the window length (500 is the typical ori length in bacteria)
	t is the threshold (for the number of appearances of a specific k-mer in a region)

*/

func FindClumps(Genome string, k, L, t int) []string {
	// fmt.Println("Finding clumps.")

	Patterns := make([]string, 0)
	n := len(Genome)

	//this is the outer loop for every window length!
	for i := 0; i < n-L+1; i++ {
		//now let's define what is within the window!
		Window := Genome[i : i+L]
		freqMap := FrequencyTable(Window, k) //this is going to add to the frequency table the appearances of any frequent k-mers

		//Pick out the values which pass the threshold
		for key, val := range freqMap {
			if val >= t {
				Patterns = append(Patterns, key)
			}
		}
	}
	//now Patterns is flushed out. Lets remove any duplicates

	Patterns = utils.RemoveDuplicatesFromArray(Patterns)

	return Patterns
}

func FrequencyTable(Window string, k int) map[string]int {
	freqTable := make(map[string]int)
	n := len(Window)

	for i := 0; i < n-k+1; i++ {
		//moving window for k-mers this time
		kmer := Window[i : i+k]
		freqTable[kmer]++
	}

	return freqTable
}
