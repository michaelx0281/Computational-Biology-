package main

import (
	"fmt"
	"math"
	// "time"
	// "log"
)

func main() {
	fmt.Println("Starting HW 0!")

	n := 0

	fmt.Println("Factor of", n, "is", Factorial(n))
	fmt.Println(Permutation(3, 2))

	fmt.Println(Power(2,3))

	fmt.Println(IsDivisible(6, 0))

	fmt.Println(SumProperDivisors(6))

	fmt.Println(FibonacciArray(5))

	array:=[]int{2, 4, 6, 8, 10}

	fmt.Println(DividesAll(array, 2))

	fmt.Println(MaxIntegerArray([]int{1,3,4,5,6,7,24,6356}))

	fmt.Println(SumIntegers(1,2,3,4,5))

	fmt.Println(GCDArray(array))

	fmt.Println(IsPerfect(28))

	x:=5
	fmt.Println(NthPerfectNumber(x))

	y:=6
	fmt.Println(NextPerfectNumber(y))

	fmt.Println(ListPrimes(50))
	fmt.Println(ListMersennePrimesBelowN(50))

	fmt.Println(ListMersennePrimes(7))

	fmt.Println(NextTwinPrimes(1000))
}

//2.1 --> Build main function and subroutines for Permutations function!
/* 
	The number of ways to order n objects i n!, 
	since we have n choices for the object in position 1,  n-1 choices for the object in position 2, and so on. 
	If we only want to order k of the objects, we still have n choices for the object in position 1, n - 1 choices for the object
	in position 2, but this stops when we have the object in positiion k, where we have n - k + 1 choices. This is called the 
	permutation statistic P(n, k) and is equal to the product n*(n-1)...(n-k+1). Note that this expression can salso be written using
	factorials as n!/((n-k)!)

	Steps:
		1) Copy over the factorial function or reimplement (rem that n! = n*(n-1)!
		2) Implement Permutation following the permutation formula
*/ 

//Factorial Subroutine Function
// takes in int n and returns int n!

func Factorial(n int) int { // this is what I implemented off the top of my head --> revisit for better / more elegent solutions

	if n < 0 {
		panic("Error: negative n value cannot be processed")
	}

	store := 1
	for i:=n ; i>0; i-- {	
		store*= i
	}

	return store
}

//Permutation takes integer n and k and returns an integer representing the possible permutations with order of k
func Permutation(n, k int) int {
	return Factorial(n) / Factorial(n-k)
}

//Combination
func Combintation(n, k int) int {
	return Permutation(n, k) / Factorial(k)
}

// 2.2 More Practice with Integer Functions 

//Power
func Power(a, b int) int {
	 store := 1
	i:=0
	for i<b {
		store *= a
		i++
	}

	return store
}

//SumProperDivisors 

func SumProperDivisors(n int) int {

	sum := 0
	for i:= n-1; i>=1; i-- {
		if IsDivisible(n, i) {
			sum+=i
		}
	}

	return sum
}

func IsDivisible(x, y int) bool {
	if y == 0 {
		return false
	}
	return x % y == 0
}

// 2.3 Working with Arrays and Variadic Functions
// FibonacciArray accepts an integer n as input and returns an array of 
// length n+1 whose k-th element is the kth Fibonacci number!

func FibonacciArray(n int) []int {
	array := make([]int, 1)
	array[0] = 1
	if(n == 0) {
		return array
	} else if (n == 1) {
		array = append(array, 1)
		return array
	} else {
		array = append(array, 1)
		for i := 2; i <= n; i++ {
			array = append(array, array[i-1] + array [i-2])
		}
		return array
	}
}

//DividesAll accepts as input an array of integers a and an integer d and returns
//true if d is a divisor of every integer in a, and false otherwise

func DividesAll(a []int, d int) bool {
	for i := range a{
		if !IsDivisible(a[i], d) {
			return false
		}
	}
	return true
}

//MaxIntegerArray accepts as input an array of integers list and returns
//the maximum value of all of the integers in the array
func MaxIntegerArray(list []int) int {

	var updatedPhase int
	for incrementingPhase, value := range list {
		if incrementingPhase == 0 || value > updatedPhase {
			updatedPhase = value
		}
	}

	return updatedPhase
}

//MaxIntegers is a variadic function that is built upon MaxIntegerArray
func MaxIntegers(numbers...int) int {
	return MaxIntegerArray(numbers)
}

//SumIntegers is an arbitrary collection o fintegers numbers
//and returns the sum of all of the integers
func SumIntegers(numbers...int) int {

	sum := 0

	for i := range numbers {
		sum += numbers[i]
	}

	return sum
}

//GCD Array accepts as input an array of integers a and returns the 
//greatest common divisor of all of theintege4rs in the array. You may want to use 
//MinIntegerArray as a subroutine

func GCDArray(a []int) int {
	for i:= MaxIntegerArray(a); i>=1; i-- {
		if(DividesAll(a, i)) {
			return i;
		}
	}
	panic("Error: No GCD found")
	return 1
}

//2.4 Perfect numbers and Mersenne Primes
/* 
	A perfect number is an integer n that is equal to the sum of its 
	proper divisors(recall that proper divisors of an integer n are those small than n).
	For example, 6 is perfect because 1 + 2 + 3 = 6 and 28 is perfect because 1 + 2 + 4 + 7 + 14 = 28.
	Perfect numbers are far rarer than prime numbers; the Greeks only knew of these two as well as 496 and 8128, and only
	just over 50 pefect numbers have ever been discovered. Surely if we know of se few perfect numbers, they must be finite?
*/

//IsPerfect accepts as input an integer n and returns a boolean indicated whether it is perfect or not

func IsPerfect(n int) bool{

	sum := 0

	for i := n; i >= 1; i-- {
		if IsProperDivisor(n, i) {
			sum += i
		}
	}

	return sum == n
}

func IsProperDivisor(n, d int) bool {
	if d >= n {
		return false
	}
	return n%d == 0
}

//NextPerfectNumber accepts as input an integer n and returns 
//the smallest perfect number that is larger than n

/* 
	Utilize this pattern to find an algorithm! 

	1)		6 = 2^1(2^2 − 1)
	2)		28 = 2^2(2^3 − 1)
	3)		496 = 2^4(2^5 − 1)
	4)		8128 = 2^6(2^7 − 1)
*/

func NextPerfectNumber(n int) int {

	i := 0

	array := make([]int, 1) 
	array[0] = -1

	if n < 0 {
		n = 0
	}

	for true {
		if n >= array[i] {
			array = append(array, NthPerfectNumber(i+1))
			i++
		} else {
			return array[i]
		}
	}

	return 0;
}

//finds the nth perfect number following an algorithm
func NthPerfectNumber(n int) int {
	var powerIncrement int

	if n == 0 {
		return 0
	} else if n == 1 {
		powerIncrement = 0
	} else if n == 2 {
		powerIncrement = 1
	} else {
		powerIncrement = 1
		powerIncrement += 2*(n-2)
	}

	a := Power(2, 1+powerIncrement)
	b := Power(2, 2+powerIncrement) - 1

	return a*b
}

//MersennePrimes 
/* 
	The ancient Greeks knew that numbers of the form 2m − 1 are more likely than others to be prime; 
	such prime numbers are called Mersenne primes. When m = 4, 2m − 1 is 15, which is certainly not prime. 
	But 22−1 = 3, 23 − 1 = 7, 25 − 1 = 31, and 27 − 1 = 127 are all prime.
	
	The larger the value of n, the less the likelihood that n will be prime. 
	So when we are looking for large prime numbers, chances are that choosing a 
	random number to test for primality will be composite. It turns out that 2m − 1 
	is composite if m is composite, and that 2m − 1 has a better than typical chance of 
	being prime if m is prime. For this reasons, when trying to find the largest known prime number, 
	mathematicians test numbers of the form 2m − 1.
*/

//ListMersennePrimes accepts a integer n as input and returns 
//an array of all primes of the form 2^p-1, where p is a positive integer
//that is less than or equal to n

//I have decided that I will just copy and paste relevant functions over
//because importing seems like a hassle
func ListMersennePrimes(n int) []int {
	primes := make([]int, 0)
	mersennePrimes := make([]int, 0)
	for i:=1; i<=n; i++{
		primes = append(primes, Power(2, i) - 1)
	}

	for i:= range primes {
		if IsPrime(primes[i]) {
			mersennePrimes = append(mersennePrimes, primes[i])
		}
	}

	return mersennePrimes
}

func IsPrime(p int) bool {
	if(p == 0 || p == 1) {
		return false
	}
	for i := 2; float64(i) <= math.Sqrt(float64(p)); i++ {
		if p%i == 0{
			return false
		}
		`	`
	}

	return true
}

func ListMersennePrimesBelowN(n int) []int {
	primes := ListPrimes(n)

	mersennePrimes := make([]int, 0)
	for i:= range primes{
		primes[i] -= 1

		mod := primes[i]

		for mod % 2 == 0 {
			mod /=2
			if mod == 1 {
				mersennePrimes = append(mersennePrimes, primes[i]+1)
			}
		} 
	}
	return mersennePrimes
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

//2.5 Twin primes

/*
	Twin primes are pairs of prime numbers that are only 2 apart (such as 3 and 5, or 29 and 31). 
	Much like perfect and amicable numbers, no one knows if there are infinitely many pairs of twin primes, 
	although computations have indicated that they do go on forever, leading to the Twin Primes Conjecture 
	that there are infinitely many such pairs.
*/

//NextTwinPrimes accepts as input an integer n and returns
//the smallest pair of twin primes that are both larger than n

func NextTwinPrimes(n int) (int, int) {

	var a, b int 
	for i:=n+1; !(a>n && b>n) ; i++ {
		if IsPrime(i) {
			if IsPrime(i+2) {
				a = i
				b = i+2
			}
		}
	}

	return a, b
}