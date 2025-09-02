package main

import (
	"fmt"

	. "github.com/michaelx0281/Computational-Biology/src/research_2025/eutils" //dealing with submodules is such a pain...look to do it later if you want..the github submodule already took me substantial time, don't want to do the same for go modules any honestly
)

func main() {
	str := Fcgi(ESearch, Gene, Term("Escherichia coli")).Assemble()
	fmt.Println(str)

	// fmt.Println(SpliceInsert("Hello, my name is Michaela!"))
}
