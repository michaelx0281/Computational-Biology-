package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("Welcome to Problem Set 2")

	fmt.Println(generateNumberInRange(4, 5))
}

// 2.1
func WeightedDie() int {
	//3 is 0.5, everything else is 0.1 probability
	index := rand.Float64()

	if index >= 0.5 {
		return 3
	} else if index <= 0.1 {
		return 1
	} else if index <= 0.2 {
		return 2
	} else if index <= 0.3 {
		return 4
	} else if index <= 0.4 {
		return 5
	} else if index < 0.5 {
		return 6
	}
	return 3
}

func EstimatePi(numPoints int) float64 {
	//for each point there is an ordered pair (a, b)
	//for each point, general a range of -1 to 1 for a and -1 to 1 for b

	inCircleCounter := 0
	for i := 0; i < numPoints; i++ {
		if pointInCircle() {
			inCircleCounter++
		}
	}

	return float64(inCircleCounter) / float64(numPoints) * 4.0
}

func pointInCircle() bool {
	a := rand.Float64()
	b := rand.Float64()

	//now determine if the point is in the circle --> of which the probability approaches the limit of pi/4
	// x^2 + y^2 = 1 is the equation of our circle

	if math.Pow(a, 2)+math.Pow(b, 2) <= 1 {
		//now you know that the point is in the circle!
		return true
	}
	return false
}

//6.2
/*
	Two integers are called relatively prime if they do not share any divisors other than 1 (or also called coprime!)
*/

func RelativelyPrime(a, b int) bool {
	return EuclidGCD(a, b) == 1
}

func EuclidGCD(a, b int) int {
	for a != b {
		if a > b {
			a = a - b
		} else {
			b = b - a
		}
	}
	return a
}

func RelativelyPrimeProbability(lowerBound, upperBound, numPairs int) float64 {
	primeCounter := 0

	for i := 0; i < numPairs; i++ {
		if pairIsPrime(lowerBound, upperBound) {
			primeCounter++
		}
	}

	return float64(primeCounter) / float64(numPairs)
}

func pairIsPrime(lowerBound, upperBound int) bool {
	if upperBound < lowerBound {
		panic("Upperbound should not be greated than lowerbound.")
	}
	a := generateNumberInRange(lowerBound, upperBound)
	b := generateNumberInRange(lowerBound, upperBound)

	return RelativelyPrime(a, b)
}

func generateNumberInRange(lowerBound, upperBound int) int {
	x := rand.Intn(upperBound + 1 - lowerBound)
	return lowerBound + x
}
