package main

import (
	"os" //for reading from data
	"strconv"
	"strings"
)

// ReadElectoralVotes processes the number of electoral votes for each state.
// Input: a filename string.
// Output: a map that associates each state name (string) to an unsigned integer corresponding to its number of Electoral College Votes
func ReadElectoralVotes(filename string) map[string]uint {
	electoralVotes := make(map[string]uint)

	//read in the file contents.
	fileContents, err := os.ReadFile(filename)
	//fileContents is a slice of bytes (ordered)

	Check(err)

	giantString := string(fileContents)

	//let's split the string into lines
	//look for occurences of a 'new line' symbol
	lines := strings.Split(giantString, "\n") //strings.Split() takes in a string and the symbol to look for when cutting and returns a slice of strings

	//range over lines, parse each line, and add values to our map
	for _, currentLine := range lines {
		//statenames are always separated by a comma
		lineElements := strings.Split(currentLine, ",") // spliting via the comma
		//lineElements has two items: the state name and the number of electoral votes (as a string)
		stateName := lineElements[0]

		//parse the number of electoral votes
		numVotes, err := strconv.Atoi(lineElements[1])
		Check(err)

		//convert int to uint and place into the map
		electoralVotes[stateName] = uint(numVotes)
	}

	return electoralVotes
}

// ReadPollingData parses polling percentages from a file.
// Input: a filename string
// Output: a map of state names (strings) to floats corresponding to the percentages for candidate 1
func ReadPollingData(filename string) map[string]float64 {
	candidate1percentages := make(map[string]float64)

	fileContents, err := os.ReadFile(filename)

	Check(err)

	giantString := string(fileContents)

	lines := strings.Split(giantString, "\n")

	//range over each line of the file and parse the data
	for _, currentLine := range lines {
		//split the current line into each of its elements
		lineElements := strings.Split(currentLine, ",")

		//lineElements has 3 things, the state names and the two percentages
		stateName := lineElements[0]
		percentage1, err := strconv.ParseFloat(lineElements[1], 64)
		Check(err)

		//normalize the percentage (divide by 100) and set the appropiate map value
		candidate1percentages[stateName] = percentage1 / 100.0
	}

	return candidate1percentages
}

// Check takes as input a variable of type error.
// If the variabe is not nil it panics
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

//this is a good one to let ai handle because it is incredibly boring (that's what Dr. Compeau said at the very least..)
