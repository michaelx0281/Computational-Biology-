package main

import (
	"github.com/michaelx0281/Computational-Biology/src/utils"
)

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

//remember that this is a really slow implementation! Let's optimize this!
/*

The steps in approaching this is to:
Make it such that there is only one freqTable created instead of a bajillion somewhere
Make it such that the window doesn't shifts over by changing indices instead of making the program check for every kmer in newest window
--> just add the newest kmer

A []
B  []
Between A and B, one kmer got deleted, and a new one got created in the latest window within length n of the entire genome

--> only add the last one!
*/

func FindClumpsOptimized(Genome string, k, L, t int) []string {
	// fmt.Println("Finding clumps.")

	Patterns := make([]string, 0)
	n := len(Genome)

	i := 0
	Window := Genome[i : i+L]

	freqTable := FrequencyTable(Window, k)

	//the above covered everything within the first window!

	for i < n-L {
		//append strings that pass the threshold here!
		for key, val := range freqTable {
			if val >= t {
				Patterns = append(Patterns, key)
			}
		}

		//freqTable[Window[L-k+1:]+string(Genome[i+L])]++
		//you have to remove the first one!
		// fmt.Println("old k" + Window[:k-1])
		freqTable[Window[:k]]--
		if freqTable[Window[:k]] == 0 {
			delete(freqTable, Window[:k])
		}
		//iterate
		Window = Window[1:] + string(Genome[i+L]) //concatenation has a bad rep, but over here, the length isn't going to get any longer, so it should be fine

		//you forgot to update the actual freqTable you dummass
		freqTable[Window[L-k:]]++ // it's not L-k+1 here!
		// fmt.Println("end k", Window[L-k:])
		i++
	}

	//grab the last value
	for key, val := range freqTable {
		if val >= t {
			Patterns = append(Patterns, key)
		}
	}

	//now Patterns is flushed out. Lets remove any duplicates

	Patterns = utils.RemoveDuplicatesFromArray(Patterns)

	return Patterns
} // YAY it WORKS! And it isn't too slow. This is essentially just review of what I did before, however, without guidance on the optimization steps and self implementation! It sure took a longer time that I really wanted to spend on it!

/*

Some mistakes that I made:

n - k + 1 instead of n-k (because I started with 0 outside of the loop and started to increment right afterwards, everything is shifted over by one)

Not updating the freqTable
Updating only the new kmer values without removing old kmer values
Removing old kmer values after adding new values (should be reverse and before incrementing window to make it easier!)
Only fully delete the old kmer if the value in the freqTable == 0, otherwise just --.

*/
