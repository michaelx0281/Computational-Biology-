package main

// import "path/filepath"

// import "fmt"

// GetExtendedNeighbors takes in a pattern (read), the overlap graph and maxK.
// It returns the extendedNeighbors list. For each neighbor *n* in this list,
// distance between n and pattern is between 2 to maxK.
func GetExtendedNeighbors(pattern string, adjList map[string][]string, maxK int) []string {
	var result []string
	visited := make(map[string]bool)
	visited[pattern] = true
	dfs(adjList, 0, maxK, pattern, &result, visited)
	return result
}

func dfs(adjList map[string][]string, depth int, maxK int, currentNode string, result *[]string, visited map[string]bool) {
	if depth > 1 && depth <= maxK {

		if !visited[currentNode] {

			*result = append(*result, currentNode)
			visited[currentNode] = true
		}
	}

	if depth < maxK {
		for _, neighbor := range adjList[currentNode] {
			if visited[neighbor] == false {
				dfs(adjList, depth+1, maxK, neighbor, result, visited)
			}

		}
	}
}
