package main

import "fmt"

func main() {
	fmt.Println("Loops!")

	// var count uint = 10
	// for ; count >= 0 ; count -- { 
	// --> it is an uint -->( 0 -> -1 ) --> subtracting 1 from 0 uint --> underflow into largest possible uint
	// 	fmt.Println(count)
	// } fmt.Println("Blast Off!! ")
	// try to be careful with unsigned ints :) 

	fmt.Println(Factorial(5))

	fmt.Println(Factorial(0))

	fmt.Println(Factorial(18)) // this is integer overflow!j

	fmt.Println(FindHighestFactorial()) 
}

//AnotherFactorial takes as input an integer n and returns n!
func AnotherFactorial(n int) int {
	if n < 0 {
		panic("Error: negative input given to Factorial()")
	}
	p := 1 

	for i:=1; i<=n; i++ {
		// i := 1 is called the "initialization step"
		// i <= is called the condition
		// i++ is called the post-condition (increment)

		p = p*i 

	}

	return p
}

//AnotherSum takes an integer n as input and returns the sum
// of the first n positive integers, using a for loop
func AnotherSum(n int) int {
	if n < 0 {
		panic("Error: negative input given to Factorial()")
	}

	sum := 0

	for k := 1; k<=n; k++ {
		sum += k
	}

	return sum
}

//mathematical aside about Carl Fredrick Gauss
// --> summing the first 100 positive integers
// Gauss came back with 5050 very quickly!
// fasting way of producing first n integers --> not that many additions at all!
// n(n+1)/2 --> :) 

func GaussSum(n int) int {
	return n*(n+1) / 2
} // simple, fast, and elegant solution :3
// similar thing in Euclid's GCD and other algorithms too
//two ways --> for loop vs while loop

func YetAnotherFactorial(n int) int {
	if n < 0 {
		panic("Error: negative input given to Factorial()")
	}

	p := 1

	for i := n; i >=1; i++ {
		p *= i; 
	} // same thing as the other one, js multiplying in reverse order

	return p
} // this is yet another fine way of approaching this problem!

//SumEven taks as input an integer k and returns the sum
//of all even positive integers up to and (possibly) inlcuding k.
func SumEven(n int) int {

	if n < 0 {
		panic("Error: negative input given to Factorial()")
	}

	sum := 0

	for i := 2; i <= n; i++ {
		if i % 2 == 0 {
			sum += i
		}
	}

	/* 
	for i := 2; i <= n; i+=2 {
		sum += i
	}

	*/

	return sum
}

//mathmatical fact: n! = n*(n-1)!
//so, 1 = 1*0!, and therefore 1=0!                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            

//factorial takes as input an integer n and returns n! =
//n*(n-1)*(n-1)*...*2*1
func Factorial(n int) int {
	p := 1 // will store product
	i := 1 // serves as a counter 

	if n < 0 {
		panic("Error: negative input given to Factorial()")
	}

	// Go has no keyword while and has 'for' keyword instead
	for i <= n {
		p = p*i
		i += 1
	}

	return p
}

func FindHighestFactorial() int {
	n:=0
	for n < 30 {
		if Factorial(n) <= 0 {
			return n;
		}

		n++
	}
	return 0
}

//SumFirstNIntegers takes as input an integer n
//and returns the sum of the first n positive integers.
func SumFirstNIntegers(n int) int {
	if(n < 0) {
		panic("Error: invalid negative value. Function takes in positive numbers only")
	}

	sum := 0

	i := 1

	for i <= n {
		sum += i
		i++
	}
	return sum

	// return n*(n+1)/2 --> haha ik it works but ig i shouldn't submit something like this
}