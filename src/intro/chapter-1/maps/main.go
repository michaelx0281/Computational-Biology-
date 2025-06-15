package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Maps in Go.")

	var polls map[string]float64

	polls = make(map[string]float64)

	//in practice, polls := make(map[string]float64)

	//unlike slices --> you don't need to define the length of the map in advance

	polls["Pennsylvania"] = 0.517
	polls["Ohio"] = 0.488
	polls["Texas"] = 0.378
	polls["Florida"] = 0.5

	fmt.Println("the number of states in the map is", len(polls))

	//getting rid of Florida
	delete(polls, "Florida") // give the name of the map and the key you want to delete

	//array and slice literals

	/*
		dnaAlphabet := [4]byte{'A', 'C', 'G', 'T'}
		primes := [5]int{2, 3, 5, 7, 11} // slice ==> can change length later if you want
	*/

	//map literal
	electoralVotes := map[string]uint{
		"Pennsylvania": 20,
		"Ohio":         18,
		"Texas":        38, //Go demands final comma for consistency
	}

	UpdateVotes(electoralVotes) //maps are pass by reference!

	fmt.Println("The number of electoral votes in Pennsylvania is now", electoralVotes["Pennsylvania"])

	// for state, votes := range electoralVotes { //ranging over a map is functionally random --> go does not want to establish an order itself and break people's programming ---> do it yourself!
	// 	fmt.Println("The number of votes in", state, "is", votes)
	// }

	PrintMapAlphabetical(electoralVotes)
}

// if you care about the order
func PrintMapAlphabetical(dict map[string]uint) {
	//sort the keys of the map, then range over the sorted keys to print key-value pairs
	keys := make([]string, len(dict)) //--> they are now in a slice
	i := 0

	for key := range dict {
		//grab the keys
		keys[i] = key
		i++
	}

	//sort the keys alphabetically
	sort.Strings(keys)

	//let's range over the keys and print the associated dictionary value
	for _, key := range keys {
		//get key
		fmt.Println("The value associated with", key, "is", dict[key]) // ranging over the keys is basically the same as ranging over the dictionary
	}
}

func UpdateVotes(electoralVotes map[string]uint) {
	electoralVotes["Pennsylvania"] = 19
	electoralVotes["Ohio"] = 17
	electoralVotes["Texas"] = 40
}
