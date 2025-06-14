package main

import (
	"fmt"
	"time"
	"log"
)

func main() {
	fmt.Println("Implementing two GCD algorithms")

	fmt.Println(TrivialGCD(63, 42))
	fmt.Println(EuclidGCD(63, 42))

	x := 378202680
	y := 273147943

	//Timing TrivialGCD
	start := time.Now() // starts stopwatch
	TrivialGCD(x, y)
	elapsed := time.Since(start) // stops stopwatch
	log.Printf("TrivialGCD took %s", elapsed) // c? this line looks like c. this is very strange
	// the above "prints to console in a pretty way" 

	start2 := time.Now()
	EuclidGCD(x, y)
	elapsed2 := time.Since(start2)
	log.Printf("Euclid GCD took %s", elapsed2) // it is not printing microseconds or nanoseconds --> did something happen to the library?
	//Euclid has a much deeper principle and is much shorter 
	//in implementation --> it is much better overall
}

func Min2(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

/*TrivialGCD takes as input two integers a and b and returns
their GCD, by applying a trivial algorithm attempting each
possible divisory of a and b up to their minimum.
*/
func TrivialGCD(a, b int) int {
	d := 1

	m := Min2(a, b)

	for i := 1; i <=m ; i++ {
		if a % i == 0 && b % i == 0 {
			d = i
		}
	}

	return d
}
/*EuclidGCD(a,b) takes as input two integers a and b and returns their
GCD, follwing Euclid's algorithm*/

func EuclidGCD(a, b int) int {
	for a != b {
		if a > b{
			a = a - b
		} else {
			b = b - a
		}
	}
	return a
}