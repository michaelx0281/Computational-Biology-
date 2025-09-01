package main

import (
	"fmt"
	. "research2025/eutils" //the dot makes it so that i don't need to prefix everything with 'eutil'
)

func main() {
	str := Fcgi(ESearch, Gene, Term("Escherichia coli")).Assemble()
	fmt.Println(str)

	// fmt.Println(SpliceInsert("Hello, my name is Michaela!"))
}
