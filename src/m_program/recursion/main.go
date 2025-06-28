package main

import (
	"fmt"
)

func main() {
	fmt.Println("Recursion!")
	fmt.Println(RecursiveFactorial(5))
	fmt.Println(RecursiveFibonacci(3)) // starts at the 0th index number --> //big tree bad!!!
	// fmt.Println(InverseRecursiveFactorial(120))

	for i := 0; i < 50; i++ {
		fmt.Println(RecursiveFibonacci(i))
	}
}

func RecursiveFactorial(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * RecursiveFactorial(n-1)
	}
}

func InverseRecursiveFactorial(n int) int {
	// n = n*(n-1)*(n-2)...(n-(n-1)) --> at (n-n) return 1
	// fmt.Println("n is equal to", n)
	if n == 1 {
		return 1
	} else {
		return n / InverseRecursiveFactorial(n-1)
	}
}

func RecursiveFibonacci(n int) int { //this times out after 10 seconds! --> calling this would get about 2^n!!
	//this exhausts the memory that you have to the stack
	if n == 0 || n == 1 {
		return 1
	} else {
		return RecursiveFibonacci(n-1) + RecursiveFibonacci(n-2)
	}
}

//instead of this, you can start at 0 --> use big for loop and array
// array + small numbers that builds up --> this is called Dynamic Programming!
// utilize other data structures to cut down costs of execution (time / memory)
//there is no waiting --> all of the stuff is just there

//Dynamic Programming was actually all made up (Richard Bellman, a Wise Man)

//Define s(i, j)
//start at s(0,                                                                                                                                                                                                                                                                               0)
