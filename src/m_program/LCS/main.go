package main

import "fmt"

func main() {
	fmt.Println("Adapted to a general alignment section.")

	str1 := "ATCGA"
	str2 := "ATGGA"
	profile := makeProfile(str1, str2)

	c_seq := profile.ConsesusSequence()

	fmt.Println(c_seq)
}
