package main

import (
	"sort"
)

/*

BetaDiversity matrix takes a map of frequency maps along with a distance metric
("Bray-Curtis" or "Jaccard" as input)...
It returns a slice of strings corresponding to the sorted names of the keys
in the map, aloing with along with matrix of distances whose (i,j)-th
element is the distance between the i-th and j-th samples using the input metric.
Input: a collection of frequency maps samples and a distance metric
Output: a list of sample names and a distance matrix where D_i,j is the distance between
sample_i and sample_j according to the given distance metric


/* Alpha diversity is within its own dataset */
/* Diversity in terms of types of species and amounts (has to be normalized) */

/* Beta diversity is the difference between two different lists */
/* They are different forms of looking at the population numbers of each of the species */

func BetaDiversityMatrix(allMaps map[string](map[string]int), distanceMetric string) ([]string, [][]float64) {

	var matrix [][]float64 //creates a matrix
	samples := make([]string, 0)
	/* ranging over the keys and compiling a slice of all of the sample names */
	for sample := range allMaps {
		samples = append(samples, sample)
	}

	sort.Strings(samples)
	// //distance matrix --> the Jaccard distance
	// matrix = make([][]float64, len(samples))

	// //making the defining the type of the outer matrix (the previous one defines the inner matrix only) --> wait why is the language like this (every other language doesn't do this)
	// for i := 0; i < len(samples); i++ {
	// 	matrix[i] = make([]float64, len(samples)) //here it is making a slice of float64's for each element
	// 	//what happens if you do matrix[i] = make([]int, len(samples))
	// }

	matrix = Make2D_2[float64](len(samples), len(samples))

	for i := 0; i < len(samples); i++ {
		for j := 0; j < len(samples); j++ {
			if distanceMetric == "Jaccard" {
				matrix[i][j] = JaccardDistance(allMaps[samples[i]], allMaps[samples[j]])
			} else {
				matrix[i][j] = BrayCurtisDistance(allMaps[samples[i]], allMaps[samples[j]])
			}
		}
	} // surprisingly, this is the end to everything here!

	return samples, matrix
}

func Make2D[T any](n, m int) [][]T {
	matrix := make([][]T, n)
	rows := make([]T, n*m)
	for i, startRow := 0, 0; i < n; i, startRow = i+1, startRow+m {
		endRow := startRow + m
		matrix[i] = rows[startRow:endRow:endRow]
	}
	return matrix
}

func Make2D_2[T any](n, m int) [][]T {
	matrix := make([][]T, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]T, m)
	}

	return matrix
}

//code challenge: --> turn rangeInt() into a generic type method that ranges  to a length of int n with an optional body that takes in what is to be done within the body (I am not sure if this is actually possible, but this seems pretty fun)
