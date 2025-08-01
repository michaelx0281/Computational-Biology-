package main

/*

Code Challenge: Solve the Pattern Matching Problem.

Input: Two strings, Pattern and Genome.
Output: A collection of integers specifying all starting positions where Pattern appears as a substring of Genome.

*/

func PatternMatching(Pattern, Genome string) []int {
	arr := make([]int, 0) //append to the arr gradually

	n := len(Genome)
	k := len(Pattern)

	//range a sliding window of size Pattern over Genome to find matching substrings
	for i := 0; i < n-k+1; i++ {
		window := Genome[i : i+k]

		if Pattern == window {
			//record the index and put into the array
			arr = append(arr, i)
		}
	}

	return arr
}
