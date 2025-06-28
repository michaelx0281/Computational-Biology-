package main

import (
	"fmt"
)

func main() {
	fmt.Println("Metagenomics.")

	a := Make2D_2[string](2, 2)
	c := Make2D[int](2, 2)

	a[0][0] = "abcd"
	a[0][1] = "efgd"
	a[1][0] = "hilo"
	a[1][1] = "hannah"

	b := []string{"a", "b", "c"}

	fmt.Println(a, b)
	fmt.Println(c)

	//very so duper cool demonstration thingy --> I am so amazing guys!!
	//I totally did not learn this on stack overflow
}
