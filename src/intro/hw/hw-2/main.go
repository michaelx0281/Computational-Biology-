package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("Welcome to Problem Set 2")

	a := []int{1, 2, 3, 4, -1000000, -1000000, -1000000}
	fmt.Println(ComputePeriodLength(a))
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

//6.3

func HasRepeat(a []int) bool {
	index := make(map[int]bool)
	for _, value := range a {
		if index[value] {
			return true
		}
		index[value] = true
	}

	return false
}

func SimulateOneBirthdayTrial(numPeople int) bool {
	return HasRepeat(generateBirthdays(numPeople))
}

func generateBirthdays(numPeople int) []int {
	birthdays := make([]int, numPeople)

	for i := 0; i < numPeople; i++ {
		birthdays[i] = rand.Intn(366)
	}

	return birthdays
}

func SharedBirthdayProbability(numPeople, numTrials int) float64 {

	sharesBirthday := 0
	for i := 0; i < numTrials; i++ {
		if SimulateOneBirthdayTrial(numPeople) {
			sharesBirthday++
		}
	}

	return float64(sharesBirthday) / float64(numTrials)

}

//6.4 In a state of sin with the middle-square PRNG

func CountNumDigits(x int) int {
	if x == 0 {
		return 1
	}

	if x < 0 {
		x *= -1
	}

	i := 0
	for x != 0 {
		x /= 10
		i++
	}

	return i
}

func SquareMiddle(x, numDigits int) int {
	if numDigits%2 != 0 || x < 0 || numDigits <= 0 || CountNumDigits(x) > numDigits {
		return -1
	}
	border := numDigits / 2

	//square x
	x *= x

	length := CountNumDigits(x)
	zeroes := 0
	if length < numDigits*2 {
		zeroes = numDigits*2 - length
	}
	x = x % Pow10(length-(border-zeroes))

	x /= Pow10(border)

	return x
}

func Pow10(a int) int {
	result := 1
	for i := 0; i < a; i++ {
		result *= 10
	}
	return result
}

/*
	GenerateMiddleSquareSequence(seed, numDigits)
    seq ← array of length 1
    seq[0] ← seed
    while HasRepeat(seq) is false
        seed ← SquareMiddle(seed, numDigits)
        seq ← append(seq, seed)
    return seq
*/

func GenerateMiddleSquareSequence(seed, numDigits int) []int { // I did it but by this point I have no idea what was going on..
	sequence := make([]int, 1)
	sequence[0] = seed

	for !HasRepeat(sequence) {
		seed = SquareMiddle(seed, numDigits)
		sequence = append(sequence, seed)
	}

	return sequence
}

func ComputePeriodLength(a []int) int {

	index := make(map[int]int)
	difference := 0
	if HasRepeat(a) {
		for i, value := range a {
			if index[value] != 0 {
				difference = i - index[value]
			}
			index[value] = i
		}
	}

	return difference
} //storing a map to an int where the value of the array is the key and the index of the array is value to the map is pretty cool!

//Another PRNG --> Linear congruential generators

// y = Remainder(a*x + c, m)
// y = (a*x + c)%m

func GenerateLinearCongruenceSequence(seed, a, c, m int) []int {
	sequence := make([]int, 1)

	sequence[0] = seed

	i := 1
	for !HasRepeat(sequence) {
		sequence = append(sequence, (a*sequence[i-1]+c)%m)
		i++
	}

	return sequence
}
