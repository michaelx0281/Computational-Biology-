package main

import "slices"

// GetTrimmedNeighbors takes in a string pattern (read), an adjacency list and maxK.
// It returns all n-transitivity edges in the adjList of the current read (pattern) for n <= maxK.
func GetTrimmedNeighbors(pattern string, adjList map[string][]string, maxK int) []string {
	neighbors := adjList[pattern]
	extendedNeighbors := GetExtendedNeighbors(pattern, adjList, maxK)
	numNeighbors := len(neighbors)

	for i := 0; i < numNeighbors; i++ {
		if slices.Contains(extendedNeighbors, neighbors[i]) {
			neighbors = slices.Delete(neighbors, i, i+1)
			numNeighbors--
			i--
		}
	}

	return neighbors
}
