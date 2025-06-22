package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	fmt.Println("Lets simulate an election!")

	electoralVoteFile := "data/electoralVotes.csv"
	pollFile := "data/debates.csv"

	//now, read them in and store as maps
	electoralVotes := ReadElectoralVotes(electoralVoteFile)
	polls := ReadPollingData(pollFile)

	numTrials := 1000000
	marginOfError := 0.1

	probability1, probability2, probabilityTie :=
		SimulateMultipleElections(polls, electoralVotes, numTrials, marginOfError)

	fmt.Println("Probability of candidate 1 winning:", probability1)
	fmt.Println("Probability of candidate 2 winning:", probability2)
	fmt.Println("Probability of tie:", probabilityTie)

}

// SimulateMultipleElections runs a Monte Carlo simulation with multiple trials to simulate a presidentail election.
// Input: polling data as a map of states to percentages for candidate 2, a map of state names to Electoral College votes, an integer coreresponding to the number of trals to run, and a decimal margin of error (confidence).
// Output: three probabilities corresponding to the likelihood of each of two candidates winning or a tie.
func SimulateMultipleElections(
	polls map[string]float64,
	electoralVotes map[string]uint,
	numTrials int,
	marginOfError float64) (float64, float64, float64) {

	winCount1 := 0
	winCount2 := 0
	tieCount := 0

	//simulate a single election n times and update count each time
	for i := 0; i < numTrials; i++ {
		//simulate one eletion
		votes1, votes2 := SimulateOneElection(polls, electoralVotes, marginOfError)

		if votes1 > votes2 {
			winCount1++
		} else if votes2 > votes1 {
			winCount2++
		} else {
			tieCount++
		}
	}

	//divide number of wins by number of trials
	probability1 := float64(winCount1) / float64(numTrials)
	probability2 := float64(winCount2) / float64(numTrials)
	probabilityTie := float64(tieCount) / float64(numTrials)

	return probability1, probability2, probabilityTie
}

func SimulateOneElection(
	polls map[string]float64,
	electoralVotes map[string]uint,
	marginOfError float64) (uint, uint) {

	var collegeVotes1 uint
	var collegeVotes2 uint

	//range over all of the stets, and simulate the election in each one.

	for state, pollingValue := range polls {
		//first lets grab the number of EC votes
		numVotes := uint(electoralVotes[state])

		//let's adjust the polling value with some noise
		adjustedPoll := AddNoise(pollingValue, marginOfError)

		//who won the state? (based on adjusted number)
		if adjustedPoll >= 0.5 {
			collegeVotes1 += numVotes
		} else {
			collegeVotes2 += numVotes
		} //not checking for ties
	}

	return collegeVotes1, collegeVotes2

}

// AddNoise adjusts a polling percentage based on some randomization sampled from a normal distribution
// Input: Two decimals, a polling value, and a margin of error.
// Output: A decimal corresponding to the adjusted polling value
func AddNoise(pollingValue, marginOfError float64) float64 {

	x := rand.NormFloat64()
	//x has 95% chance from being between -2 and 2 at 2 standevs

	x /= 2.0
	//x has a 95% chance of being between -1 and 1

	x *= marginOfError
	//x has a 95% chance of being between -marginOferror and +marginOfError

	//want: x to have a 95 percent chance of being between - marginOfError and + marginOfError
	return x + pollingValue
}

//Early Polls
//with 0.05 margin of error, 99 percent chance of Hillary Clinton Victory (versus DT)
//with 0.1 margin of error, 98 percent

//Convention Polls
//with 0.1 margin of error, 99%

//Debates
//with 0.1 margin of error, >99%

//Our model isn't "bad", the Monte Carlo simulation model is just bad in general (all of the models)
//predicting the future is hard
