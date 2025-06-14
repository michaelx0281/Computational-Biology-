package main

import (
	"fmt"
)

func main() {
	fmt.Println("Arrays (and slices)")

	var list [6]int  // --> [0 0 0 0 0 0] receives default values
	// GO uses 0-based indexing

	list[0] = -8
	i := 3
	list[2*i-4] = 17
	list[len(list) - 1] = 34

	//no out of bounds and no negative indices in Go

	list2 := []int{1,2,3,4,5,6} // this is a slice literal
	fmt.Println(list)
	fmt.Println(list2)

	
	// a slice has variable length
	// slice declarations are a little different
	var a []int // right now, a h as a special value nil
	n := 4

	// slices must be made
	a = make([]int, n) // there is length that could be changed

	//we set values of arrays similarly to those of slices
	a[0] = 6
	a[2] = -33

	fmt.Println(a) // this looks just like an array atp

	//one-line declarations are used in practice
	b := make([]int, n+2)
	fmt.Println(b)

	fmt.Println(FactorialArray(6))

	var c [6]int 

	d := make([]int, 6)

	ChangeFirstElementArray(c)
	ChangeFirstElementSlice(d)

	fmt.Println("C is now", c)
	fmt.Println("D is now", d)

	fmt.Println(MinIntegers(45,32,53,3))
}

//Variatic Functions take a variable number of inputs!

//MinIntegers taks as input an 
//arbitrary number of integers and returns their minimum value
func MinIntegers(numbers...int) int {
	// Go will create a slice called numbers for you

	if len(numbers) == 0 {
		panic("Error: empty slice given to MinIntegerArray")
	}

	return MinIntegerArray(numbers)
}

//MinIntegerArray takes as input a slice of integers and returns
//the minimum value in that slice.

func MinIntegerArray(list []int) int {

	if len(list) == 0 {
		panic("Error: empty sice given to MinIntegerArray")
	}

	min := list[0] // minimum value of m that would be updated

	for i := 1; i <= len(list)-1; i++ {
		if min > list[i] {
			min = list[i]
		}
	}

	/* 
	Can also use:

	for i, val := range list { //the equivalent of for i := 0; i < len(list); i++
		if i == 0 || val < min {
			min = val
		}
	}
	
	*/

	return min
}



// no outputs in these bottom two functions

/*
 is pass by reference meaning passing a 
 reference to the address and location in memory? aka using a pointer? 
 
pass by value == copied made and no original copy change
pass by reference == original copy is changed

pass by value replicates and creates a copy of value stored in memory
pass by reference directly provides pointer to original value in memory?
 
*/

func ChangeFirstElementArray(a [6]int) { 
	//6 is a constant --> we are okay here

	a[0] = 1 // pass by values --> a copy of c gets created, which is destroyed!
}

func ChangeFirstElementSlice(a []int) { // you get access to the slice itself!
	a[0] = 1 // can change d! pass by reference --> full understanding is a little more complicated
}

//Factorial --> compute all of the factorials up until n factorial

//FactorialArray taks as input an integer n, and it returns an
//array of length n+1 whose k-th element is k!
func FactorialArray(n int) []int {
	if n < 0 {
		panic("Error: negative value input given to FactorialArray")
	}

	// var fact [n+1]int // --> instead of the [n+1] use slice to avoid compiler errors

	fact := make([]int,  n+1)
		

	//recall; 0! = 1 --> this is very similar to recursion almost! --> base case to be made with k[0] = 1
	fact[0] = 1

	//range k from 1 to n, and use the fact that k! = k*(k-1)!
	for k := 1; k <= n; k++ {
		fact[k] = k * fact[k-1]
	}
	return fact
}