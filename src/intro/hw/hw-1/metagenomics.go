package main

import (
	"fmt"
)

func main() {
	fmt.Println("Metagenomics HW Assignments")

	// 	hash :=
	// 		map[string]int{
	// 			"a": 1,
	// 			"b": 2,
	// 			"c": 3,
	// 		}

	// 	fmt.Println(Pow(2, 3))
	// 	fmt.Println("The total sum is: ", SimpsonsIndex(hash))

	hash1 := map[string]int{
		"a": 2,
		"b": 4,
	}

	hash2 := map[string]int{
		"a": 4,
		"b": 3,
		"c": 2,
	}

	fmt.Println(SumOfMinima(hash1, hash2))
	fmt.Println(BrayCurtisDistance(hash1, hash2))
	fmt.Println((SumOfValues(hash1) + SumOfValues(hash2)) / 2)
}

//Richness
/*
	Input: frequency table 'hash'
	Output: number of non-zero elements 'richness'

*/
func Richness(sample map[string]int) int {
	richness := 0

	reads := make([]string, len(sample))

	for i := range sample {
		if sample[i] == 0 {
			continue
		}
		reads = append(reads, i)
		richness++
	}

	return richness
}

//if asked to calculate a collectable probability in a situation where probabilities calculated are independent, just some the probabilities up!

//I am not sure what MaxMap() function is so I will just write from scratch

func SumOfValues(sample map[string]int) int {
	sum := 0

	for _, val := range sample {
		sum += val
	}

	return sum
}

func SimpsonsIndex(sample map[string]int) float64 {
	/*
		HERE IS THE RATIONALE

		SumOfValues would provide the total (which is the denominator)

		The value accessed in each specific key would be the numerator

		The above are then squared and summed to produce the value
	*/

	sum := 0.0

	denominator := SumOfValues(sample)

	for _, value := range sample {
		sum +=
			Pow(
				float64(value)/float64(denominator),
				2,
			)
	}

	return sum
}

// fatal error encountered: integer division

func Pow(a float64, b int) float64 {
	value := 1.0

	for i := 1.0; i <= float64(b); i++ {
		value *= a
	}

	return value
}

// Part 2

func SumOfMinima(sample1, sample2 map[string]int) int {
	//assuming that the samples have the same lengths and count
	sum := 0

	var chosenSample map[string]int
	keys := make([]string, 0)

	if len(sample1) < len(sample2) {
		chosenSample = sample1
		keys = make([]string, len(chosenSample))
	} else {
		chosenSample = sample2
		keys = make([]string, len(chosenSample))
	}
	//first gather a slice of all of the keys
	i := 0

	for key := range chosenSample {
		keys[i] = key
		i++
	}

	// for key, val := range sample1 {
	// 	if val != chosenSample[key] && val < chosenSample[key] {
	// 		chosenSample[key] = val
	// 	}
	// }

	// for key, val := range sample2 {
	// 	if val != chosenSample[key] && val < chosenSample[key] {
	// 		chosenSample[key] = val
	// 	}
	// }

	for i := 0; i < len(chosenSample); i++ {
		if sample1[keys[i]] < sample2[keys[i]] {
			sum += sample1[keys[i]]
		} else {
			sum += sample2[keys[i]]
		}
	}

	return sum //note that you can use a Min2() and should probably use a Min2() function here --> however I didn't think about that at that time
}

func SumOfMaxima(sample1, sample2 map[string]int) int {
	//assuming that the samples have the same lengths and count
	sum := 0

	var chosenSample map[string]int

	if len(sample1) < len(sample2) {
		chosenSample = sample1
	} else {
		chosenSample = sample2
	}

	for key, val := range sample1 {
		if val > chosenSample[key] {
			chosenSample[key] = val
		}
	}

	for key, val := range sample2 {
		if val > chosenSample[key] {
			chosenSample[key] = val
		}
	}

	//first gather a slice of all of the keys

	keys := make([]string, len(chosenSample))

	i := 0

	for key := range chosenSample {
		keys[i] = key
		i++
	}

	for i := 0; i < len(chosenSample); i++ {
		if sample1[keys[i]] > sample2[keys[i]] {
			sum += sample1[keys[i]]
		} else {
			sum += sample2[keys[i]]
		}
	}

	return sum //note that you can use a Min2() and should probably use a Min2() function here --> however I didn't think about that at that time
}

func BrayCurtisDistance(sample1, sample2 map[string]int) float64 {
	return 1 - float64(SumOfMinima(sample1, sample2))/((float64(SumOfValues(sample1))+float64(SumOfValues(sample2)))/2)
}

func JaccardDistance(sample1, sample2 map[string]int) float64 {
	return 1 - float64(SumOfMinima(sample1, sample2))/float64(SumOfMaxima(sample1, sample2))
}
