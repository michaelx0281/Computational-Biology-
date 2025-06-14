package main

import (
	"fmt"
)

// ew why is the return type after the parens..

/* Takes two integers and returns their sum */
func SumTwoInts(a, b int) int { //--> function signature : (params and output type)
	return a + b
}

//DoubleAndDuplitcate takes as input a float64 and returns two copies
// of that variable

//AddONe takes an integer k as input and returns the value of k+1
func AddOne(m int) int {
	m = m+1 // oh there was no error bc ints default to 0
	return m
} 

// go uses 'pass-by' value
// when you call a function, a copy of a variable is made
// any changes made within a function does not necessarily change
// external global vars

//there is a way to make go use pass by reference --> see how to do that much later within the course

func DoubleAndDuplicate(x float64) (float64, float64) {
	return 2.0*x, 2.0*x
}

//Pi takes no inputs and returns the value of pi, 
//the mathematical constant

func Pi() float64 {
	return 3.14
}

// PrintHi: simply prints Hi to the console

func Hi() {
	fmt.Println("Hi")
}

func main() { // --> also a function --> package main means go is always looking for func main : --> this seems to behave similarly to a constructor almost?
	fmt.Println("functions")
	x := 3
	n := SumTwoInts(x, x) 
	fmt.Println(n)

	m := 17 
	fmt.Println(AddOne(m))
	fmt.Println(m)

} // func main could be reordered and it doesn't matter 

//subroutines (calling on function within another function, is also possible)