package main

import (
	"fmt"
)

func main() {
	a := Apple{color: "red", tetoPear: false, taste: "sour"}
	p := Pear{color: "green", tetoPear: true, taste: "sweet"}

	printFruit(a)
	printFruit(p)
}

func printFruit(f Fruit) {
	fmt.Println("Color", f.getColor()+". "+"I am", f.getTaste(), f.isTetoPear())
}
