package main

func GCSkew(genome string) int {
	var skew int

	n := len(genome)

	for i := range n + 1 { // mistake: n instead of n + 1 (range is i := 0; i < n; i++ not i <= n)
		if i == 0 {
			skew = 0
		} else {
			skew += UpdateSkew(genome, i-1)
		}
	}

	return skew
}

func GCSkewFull(genome string) []int {
	skew := make([]int, len(genome)+1)

	n := len(genome)

	for i := range n + 1 { // mistake: n instead of n + 1 (range is i := 0; i < n; i++ not i <= n)
		if i == 0 {
			skew[i] = 0
		} else {
			skew[i] = skew[i-1] + UpdateSkew(genome, i-1)
		}
	}

	return skew
}

func UpdateSkew(genome string, i int) int {
	previous := byte(genome[i])
	skew := 0

	switch previous {
	case 'G':
		skew = 1
	case 'C':
		skew = -1
	default:
		skew = 0
	}

	return skew
}

/*
NOTE:

YOu still need to work on the parsing utility to solve the practice exercise question and eventually look for the minimum skew.. although, it is true that it is technically unnecessary to write to a file to file the minimum skew. Just use the GCSkewFull() function to return an array / slice of integers and range over it to record any of the indices where it reaches a minimum (which you can use a MinInteger() function to achieve)

--> this and MaxInteger() are also utilities which you could / should create to better aid you within this processconst

For now tho, I am more interested in exploring further and moving past this section.
*/

func MinimumSkew(Genome string) []int {
	//return the region with the minimum skew

	//first get the full list of all of the Skews
	listSkews := GCSkewFull(Genome)
	minSkews := make([]int, 0)

	//now, range through and find the minimum value within the array

	min := MaxInt(listSkews)

	for _, skew := range listSkews {
		if skew < min {
			min = skew
		}
	}

	//now add the min skews to the list
	for i, value := range listSkews {
		if value == min {
			minSkews = append(minSkews, i)
		}
	}

	return minSkews
}

func MaxInt(list []int) int {
	max := 0

	for _, val := range list {
		if val > max {
			max = val
		}
	}

	return max
}
