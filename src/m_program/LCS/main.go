package main

import "fmt"

func main() {
	fmt.Println(LCSLength("GCAT", "GTAT"))
	// fmt.Println("Matrix", Make2D_2[int](3, 6))

	// 	matrix := Make2D_2[int](5, 3)

	// 	matrix[0][0] = 1
	// 	matrix[0][1] = 2
	// 	matrix[0][2] = 3

	// 	matrix[1][0] = 10
	// 	matrix[1][1] = 11
	// 	matrix[1][2] = 12

	// 	fmt.Println("Matrix hi", matrix)

	//every time that you backtrack diagonally, add to an index!
	fmt.Println(LCSPaths("ATCGTCC", "ATGTTATA", 7, 8))
}
