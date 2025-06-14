package main

import (
	"fmt"
	"math"
	"time"
	"log"
)

func main() {
	fmt.Println("Finding primes!")

	n := 10000000

	// fmt.Println(TrivialPrimeFinder(n))
	// fmt.Println(SieveOfEratosthenes(n))

	start := time.Now()
	TrivialPrimeFinder(n)
	elapsed := time.Since(start)
	log.Printf("TrivialPrimeFinder took %s", elapsed)
	SieveOfEratosthenes(n)
	start2 := time.Now()
	
	elapsed2 := time.Since(start2)
	log.Printf("SieveOfEratosthenes took %s", elapsed2)

	// fmt.Println(ListPrimes(n))
}
//Implementing two different methods of finding primes using
// a more trivial v. an ancient algorithms!

//TrivialPrimeFinder takes as input an integer n and returns a
//slice of bollean variables storing the primarlity of each
//nonnegative integer up to and including n.

func TrivialPrimeFinder(n int) []bool {
	primeBooleans := make([]bool, n+1)

	for p := 2; p <= n ; p++ {
		primeBooleans[p] = IsPrime(p)
	}

	return primeBooleans
}

//IsPrime takes as input an integer p and returns true if p is 
//prime and false otherwise

func IsPrime(p int) bool {
	for i := 2; float64(i) <= math.Sqrt(float64(p)); i++ {
		if p%i == 0{
			return false
		}
	}

	return true
}

func SieveOfEratosthenes(n int) []bool {
	primeBooleans := make([]bool, n+1)
	primeBooleans[0] = false
	primeBooleans[1] = false 
	
	//set everything to true!
	for p := 2; p<=n; p++ {
		primeBooleans[p] = true
	}

	for p := 2; float64(p) <=math.Sqrt(float64(n)); p++ {
		//is p prime? If so, cross off its multiples

		if primeBooleans[p] == true {
			primeBooleans = CrossOffMultiples(primeBooleans, p)
		}

	}

	return primeBooleans
}

//CrossOffMultiples takes as input an integer p and boolean slice primeBooleans, and returns 
//an updated slice in which all variables in the array whose indi ces are multiples of p (greater than p)
//have been set to false


//this is a faster process as it crosses off multiple numbers all at once!

func CrossOffMultiples(primeBooleans []bool, p int) []bool {
	n := len(primeBooleans) - 1

	//consider every multiple of p, starting wit 2p, and "cross it off" by setting its corresponding entry of the slice to false
	for k := 2*p; k<=n; k += p {
		primeBooleans[k] = false
	}

	return primeBooleans
}

//ListPrimes 

func ListPrimes(n int) []int {
	primeList := make([]int, 0)
	primeBooleans := SieveOfEratosthenes(n)

	for i := range primeBooleans {
		if primeBooleans[i] {
			primeList = append(primeList, i)
		}
	}
	return primeList
} 