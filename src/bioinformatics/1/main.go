package main

import (
	"fmt"
	"os"

	"github.com/michaelx0281/Computational-Biology/src/utils" //wow this was actually so annoying to figure out
)

func main() {
	fmt.Println("Introduction to Bioinformatics Algorithms: Chapter 1. Where in the Genome Does Replication Begin?")

	//Testing
	fmt.Println(ReverseComplement("ATGC"))

	//Exercises: You should probably try out all of them, even if you don't want to, to get the best learning experience out of this

	//actually i'm a little bit of a hippocrite, i do not feeling like doing MaxMap so I'm not doing it. It seems about the same as what I did for frequent words anyways...

	//Now onto a real Exercise Break
	PatternMatchingVibrioCholerae()
}

//Take Vibrio cholerae genome from os and print list of starting indices of any matches in a space separated list

//First time working on this, I printed out the physical sequence by accident

//Now implement the actual pattern matching before printing!

// ATGATCAAG/CTTGATCAT is a hidden message for the *ori* region of Vibrio cholerae! however, it is not the hidden box for all DnaA proteins in all bacteria!

//For example, the above pattern does not work for Thermotoga petrophila -> which thrives in extremely hot environments; its name derives from its discovery in the water beneath oil reservoirs, where temperatures can exceed 80* Celcius //still need to figure out how to type the degree symbol smh

/*

Application of the Frequent Words Problem to the ori region above reveals that the following six 9-mers appear in this region 3 or more times:

AACCTACCA   AAACCTACC   ACCTACCAC
CCTACCACC   GGTAGGTTT   TGGTAGGTT

*/

func PatternMatchingVibrioCholerae() {

	p := "CTTGATCAT" //you can also soft code this to pass in a specific pattern, but I am too lazy rn

	//TODO

	genome, err := os.ReadFile("data/Vibrio_cholerae.txt")

	utils.CheckError(err)

	//data is current a slice of bytes (byte[])

	//Let's Pattern Match!
	dataList := PatternMatching(p, string(genome))

	//Now, we just need to print it out and add spaces
	utils.PrintListSpaceSeparated(dataList) //once this works, let's write to a txt file instead!

}
