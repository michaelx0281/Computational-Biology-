package main

import (
	"fmt"
	"strconv" //string convert --> this helps bypass the ASCII
)

func main() {
	fmt.Println("Strings.")

	fmt.Println(string('A'))
	fmt.Println(string(45)) // a dash symbol '--' is printed instead --> this is in ASCII --> ASCII to string

	fmt.Println(strconv.Itoa(45)) // integer to ascii
	
	j, err:= strconv.Atoi("37") 
	// ascii to integer --> the function returns 2 values (j, err) --> without err there is compiler issue
	// for .Atoi() 
	// "[#]" e.g. "37" is correct syntax and would print 37
	// "[string]" e.g. "hi" is incorrect syntax and would cause err to return non-nil --> parsing would fail here
	
	//the conversion is OK when error variable is equal to nil. --> more in OOP
	if err != nil {
		//a problem happened
		panic(err) // catching error due to non-nil value
	}
	
	fmt.Println(j)

	pi, err2 := strconv.ParseFloat("3.14", 64) // function takes in 2 params, 
	//the representation of your var in string form and precision ->> (64 for String 64)

	if err2 != nil {
		//a problem happened
		panic(err2)
	} 

	fmt.Println("The Value of pi is", pi) // pi is 3.14 is printed --> this seems quite different from c, actually. 
	//--> use % format specifier. Java automatically handles this too, additionally

	//strconv is also very relevant and important for the task of 
	//grabbing values / variables fromt the command-line as inputs

	//talking about initialization as well as concatenation in the section below
	var s string = "Hi"
	t := "lovers"

	//concatenate these strings with + operation (which is built in!) -- similar to java. this does NOT work in c
	u := s+t
	fmt.Println(u)

	//in Go, think of strings as arrays of symbols (which are in bytes)j

	fmt.Println("The first symbol of u is", u[0]) // this is in ASCII, again!

	fmt.Println("The first symbol of u is", string(u[0]))
	fmt.Println("The final symbol of u is", string(u[len(u)-1])) // everything here is working well now!

	if t[2] == 'v' { // this works!! --> if you change it to a capital v then this does not work --> which is similar to any other language
		fmt.Println("The symbol at postition 2 of t is v.")
	}

	//This was the basics of strings in Goj


	dna := "ATGCAGT"
	fmt.Println(ReverseComplement(dna))
}

//ReverseComplement takes as input a string pattern of DNA symbols.
//It returns the reverse complement of the string
func ReverseComplement(pattern string) string {
	return Reverse(Complement(pattern))
}

//Complement takes as input a string pattern of DNA symbols.
//It returns the string formed by complementing each position of 
//the input string ('A' <--> 'T', 'C' <--> 'G')
func Complement(dna string) string { // using the below to examplify switch statements (which are also used in Java!)

	//get around this by creating a new string called dna2
	// var dna2 string // or dna2 := "" // --> this is the default value
	dna2 := make([]byte, len(dna))
	for i, symbol := range dna { // You can redeclare entire strings, just not individual symbols --> if you don't want to use i, replace with '_'
		switch symbol {
			case 'A': // Technically strings are read-only slices of symbols (bytes) --> you can't go editing individual symbols --> need to js create a new string
				dna2[i] = 'T'
			case 'C':
				dna2[i] = 'G'
			case 'G':
				dna2[i] = 'C'
			case 'T':
				dna2[i] = 'A'
			default:
				panic("Invalid sysmbol given to Complement()") // string concatenations are inefficient bc there are two many copies, especially when strings get bigger, turn into slices of bytes instead
		} // main takeways: --> change individual values then make a slice of bytes. strings are read-only slices of bytes
	}
	return string(dna2) // this converts the slice of bytes into a string
}

//Reverse takes as input a string pattern.
//It returns the string formed by reversing the positions of all
//symbols in pattern
func Reverse(pattern string) string {
	reversedMap := make([]byte, len(pattern))

	for i := len(pattern) - 1; i >=0; i-- {
		reversedMap[len(pattern) - 1 - i] = pattern[i]
	}

	return string(reversedMap)
	
}

/*
I see something pretty cool in the top right of the Code-Along Video 

"Hunting for Hidden Messages in Bacterial Genomes"

Algorithms to implement:

ReverseComplement(pattern)
	return Reverse(Complement(pattern))

Reversing and Complementing are two separate tasks
	--> this is why they are depicted as two functions (subroutines) here in this case

The model that was intitially used within the video was the way to determine the antiparallel complementary
sequence of the paired strand of DNA (in terms of nucleotide base pairs)
*/