package main

import (
	"fmt"
)


func main() {
	fmt.Println("The minimum of 3 and 4 is", Min2(4,3))

	fmt.Println(WhichIsGreater(3,5))
	fmt.Println(WhichIsGreater(42, 42))
	fmt.Println(WhichIsGreater(-2, -7))

	//returns -1, 0, 1 --> the values align and the function is working!
}

//Min2 takes two integers as input 
//and returns the minimum value out of the two
func Min2(a, b int) int {
	if a < b {
		return a
	}
	// b must be smaller (or they are equal)
	return b
}

//WhichIsGreater compares two input integers and returns 1  if
//the first input is larger, -1 if the second input is larger,
//and 0 if they are equal

func WhichIsGreater(x, y int) int {
	//we need three cases
	if x == y {
		return 0 
	} else if x > y {
		return 1
	} else {
		return -1
	}
}

//PositiveDifference takes as input two integers
//It returns the absolute value of the difference between
//these integers.

func PositiveDifference(a,b int) int {
	var c int
	if a == b {
		return 0
	} else if a > b {
		c = a - b
	} else {
		c = b - a
	} 
	return c;
}

// takes two integers and returns a boolean value
// true if same sign, false if different sign
func SameSign(x, y int) bool {
	// if( x*y < 0) { //  x and y is the same sign when their product is >0 (positive)
	// 	return false
	// } return true

	return x*y >= 0 
}

//index of comparison operators
/*

> : more than
< : less tabn
>= : greater or equal to 
<= : less than or equal to
== : equal to
!= : not equal to
! : "not", for example, if x is Boolean, then !x is false when x is true 
and true when x is false

*/

