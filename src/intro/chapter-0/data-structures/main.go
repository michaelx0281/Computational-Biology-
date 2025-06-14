package main

import "fmt"

func main() {

	var j int = 14 				//default = 0
	var x float64 = -2.3		// default = 0.0
	var yo string = "hi" 		// default => empty string ('zero string')
	var u uint = 14				// default = 0
	var symbol byte = 'H' 		//default = 0 --> 'h' : 72 --> seems like ASKII
	var statement bool = true 	// default = false

	/*shorthand declarations avoid var j int statements*/

	i := -6 					// i's type is then inferred in this instance
	hi := "yo" 					// has type string
	k := 34 					// int or uint? --> this would be an int --> can later be updated
	y := 3.7 					// type float64
	secondStatement := true 	// has type bool


	
	
	// go does not like that are there variabels declared but unused
	// compiler is a 'grammar checker'
	// complite frequently (just like pulling from git frequently)
 
	fmt.Println("Variables and types.")
	fmt.Println(j)
	fmt.Println(x)
	fmt.Println(yo)
	fmt.Println(u)
	fmt.Println( symbol) 
	fmt.Println(statement)
	
	fmt.Println(i, hi, k, y, secondStatement) // this adds a space --> slightly different from concatenation within java

	fmt.Println( 2*(i+5) * k)
	fmt.Println( 2*y - 3.16)

	fmt.Println(float64(k)*y) // --> oh it looks like there is some floating point rounding error --> also why does every language type cast inverted when compared to c...

	fmt.Println(float64(k)/14) // dividing 2 integers --> integer division --> has a natural floor function (rids itself of the remainder)

	//go doesn't allow bool(0)

	var p int = -187 // compiler error? --> starting at 0, must mean maximum - 187 --> overflow
	var s uint = uint(p)
	fmt.Println(s)

	m := 9223372036854775807 // --> why does this add one but turn negative? 

	// integer overflow --> there is only a finite amount of space that could be stored
	// something about powers of two?

	fmt.Println(m + 2)
}

//leaving small comment that my intellisense for go is not 
//autofilling in fmt --> I am not sure if it has anything to do with
//go modules being turnt off (I found online gopls works best with modules)

/* 
I think that this would work quite well for me. it is pretty 
similar to java, other than the fact that there are no semi-colons (so far anyways)
*/

//I wonder if my go extension is broken...