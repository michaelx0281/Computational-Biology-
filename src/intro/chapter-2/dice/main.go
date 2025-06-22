package main

import (
	"fmt"
	"math/rand"
	// "time"
)

func main() {
	// rand.Seed(1) // seeding with 1 --> this allows control over a process  (rand.Seed(1) used to to be default behavior)
	// rand.Seed(time.Now().UnixNano()) // use this to get different results --> this used to be something people had to do
	fmt.Println("Rolling dice.")

	// fmt.Println(rand.Int())     // print integer
	// fmt.Println(rand.Intn(10))  // print integer between 0 and 9
	// fmt.Println(rand.Float64()) // print decimal in range [0, 1) --> something different happens internally every time

	// //for years, running the code multiple times would return the same output --> the seeding of randomness is not infact truely random, however, there is a different mechanism internally now that works better at creating random results

	// fmt.Println(SumDice(2))

	numTrials := 10000000
	fmt.Println("Estimated house edge with", numTrials, "trials is:", ComputeCrapsHouseEdge(numTrials)) // Craps has a high variance!
}

//RollDie
//Input: none
//Output: a pseudorandom integer between 1 and 6, inclusively

func RollDie() int {
	return rand.Intn(6) + 1 //Why not just use rand.Intn(7) ???
}

// SumTwoDice
// Input: none
// Output: the simulated sum of two dice (bewtween 2 and 12)
func SumTwoDice() int {
	// you can take the probability of rolling each number

	/*
		roll := rand.Float64() //if this is the probability, you can divide the other ones into differently sized packets
		if roll <1.0/36.0 {
			return 2
		} else if roll < 3.0/36.0 {// we know it is bigger that 1/36 once we are here (the interval is of with 2)
			return 3
		} else if roll <6.0/36.0 {
			return 4
		} // etc.
	*/

	return RollDie() + RollDie()

}

// PlayCrapsOnce simulates one game of craps
// Input: none
// Output: true if the game is a win and false if it is a loss (from the perspective of the player)
func PlayCrapsOnce() bool {
	firstRoll := SumDice(2)
	if firstRoll == 7 || firstRoll == 11 {
		//winner!
		return true
	} else if firstRoll == 2 || firstRoll == 3 || firstRoll == 12 {
		return false //loser! :(
	} else {
		// keey rolling until you hit a 7 or your original roll
		for { //while forever
			newRoll := SumDice(2)
			if newRoll == firstRoll {
				return true //winner!
			} else if newRoll == 7 {
				return false // loser! :(
			}
		}
	}
}

// ComputeCrapsHouseEdge estimates the "house edge" of craps over multiple simulations
// Input: an integer corresponding to the number of simulations
// Output: house edge of craps (average amount of won or loss over the number of simulations)
func ComputeCrapsHouseEdge(numTrials int) float64 {
	count := 0 // will keep track of amount won (+) or lost (-)

	// run n trails and update count accordingly
	for i := 0; i < numTrials; i++ {
		outcome := PlayCrapsOnce()

		if outcome {
			count++
		} else {
			count--
		}
	}

	return float64(count) / float64(numTrials)
}

// SumDice
// Input: an integer numDice
// Output: the sum of numDice simulated dice
func SumDice(numDice int) int {
	sum := 0

	for i := 0; i < numDice; i++ {
		sum += RollDie()
	}

	return sum
}

//Simulating craps with Monte Carlo simulation (runs millions of randomized trials to simulate losses at a casino)
