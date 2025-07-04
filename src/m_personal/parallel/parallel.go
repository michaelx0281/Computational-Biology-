package main

import (
	"fmt"
	"math/rand"
	"time"

	// "log"
	"runtime"
	// "time"
)

func main() {
	fmt.Println("Parallel programming and concurrency.")

	fmt.Println("This computer has", runtime.NumCPU(), "cores available") //all 16 cores used by default

	//channel = telephone line, sending information back and forth

	c := make(chan string) //each of the cables have to have a type
	// c <- "Hello!"          // sending

	//this BLOCKS, meaning it doesn't any code in this function
	//until someone pickes up the phone somewhere
	go SayHi(c) //this calls the function

	//channels and go routines go hand-in-hand!!

	msg := <-c

	fmt.Println(msg) // receiving from the channel (pulling out)

	// this channel was synchronous! Sending and receiving get coordinated to happen at the same time!
	// you do not send a message until you pick up!

	n := 4
	// start := time.Now()
	// Factorial(n)
	// elapsed := time.Since(start)
	// log.Printf("Multiple processors took %s", elapsed)

	// runtime.GOMAXPROCS(1)

	// start2 := time.Now()
	// Factorial(n)
	// elapsed2 := time.Since(start2)
	// log.Printf("One processor took %s", elapsed2)

	// go PrintFactorials(10)  //these never had the chance to finish

	// go PrintFactorials(20)

	c2 := make(chan int)
	go Perm(1, n/2, c2) //you can't carry on if you have to wait on these
	go Perm(n/2, n, c2)

	fact1 := <-c2
	fact2 := <-c2

	//because of blocking, this wouldn't continue until both have recieved their message

	fmt.Println("n! is", fact1*fact2)

	fmt.Println("Program finished")

	fmt.Println("Running serial simulations.")
	start := time.Now()
	fmt.Println(CrapsHouseEdgeMultiProc(100000000, 1))
	elapsed := time.Since(start)

	fmt.Printf("1 procs took %s", elapsed)

	fmt.Println("Craps in parallel")
	start2 := time.Now()
	fmt.Println(CrapsHouseEdgeMultiProc(100000000, 1))
	elapsed2 := time.Since(start2)

	fmt.Printf("16 procs took %s", elapsed2)

}

func CrapsHouseEdgeMultiProc(numTrials int, numProcs int) float64 {
	count := 0 // represent my winnings (or losing)

	c := make(chan int) // you all call me how much you lost

	// play the game in parallel
	// divide the work into equal pieces and have numTrials total trials

	for i := 0; i < numProcs-1; i++ {
		currentNumbersOfTrials := numTrials / numProcs
		go TotalWinOneProc(currentNumbersOfTrials, c)
	}

	// one more goroutine corresponding to all the remaining trials
	go TotalWinOneProc(numTrials/numProcs+numTrials%numProcs, c)

	// each is going to have some amount to say home much they lost

	//call me and tell me you losses
	for i := 0; i < numProcs; i++ {
		yourLosings := <-c
		count += yourLosings
	}

	return float64(count) / float64(numTrials)
}

// TotalWinOneProc
// Input: a number of trails and a channel
// Output: nothing, you play this many times and send the message of your winnings into the channel (the 'telephone')
func TotalWinOneProc(yourNumberOfTrials int, c chan int) {
	count := 0
	// play the game appropiate number of times
	// keep track of winnings

	for i := 0; i < yourNumberOfTrials; i++ {
		outcome := PlayCrapsOnce()

		if outcome {
			count++
		} else {
			count--
		}
	}

	c <- count
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

func RollDie() int {
	return rand.Intn(6) + 1 //Why not just use rand.Intn(7) ???
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

func SayHi(c chan string) {
	c <- "Hello!"

	// this will wait until someone picks up!
}

func Factorial(n int) int {
	p := 1
	for k := 2; k <= n; k++ {
		p *= k
	}

	return p
}

func PrintFactorials(n int) {
	p := 1
	for i := 1; i < n; i++ {
		fmt.Println(p)
		p *= i
	}
}

func Perm(k, n int, c chan int) {
	//multiply the numbers from k up to but not including n
	p := 1
	for i := k; i < n; i++ {
		p *= i
	}

	c <- p // send the message taht I'm done!
}
