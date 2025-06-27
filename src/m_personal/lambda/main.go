package main

import (
	"fmt"
)

//Here is an overview of how lambda expressions work within go, taken from google and it's silly ai!

func main() {
	fmt.Println("Lambda.")

	//Define and execute an anonymous function immediately! :P
	func() {
		fmt.Println("Hello there :p")
	}() //the second paren down here is essential for things that you would like to immediately invoke

	// Assign an anonymous function to a variable
	yell := func(message string) {
		fmt.Println(message)
	}

	// Call the anonymous function through the variable
	yell("I am so amazing and I love computing!!") // this makes for great readability, even if the anonymous function stuff is pretty janky

	// Anonymous functino used as a closure

	counter := createCounter()
	fmt.Println(counter()) // 1
	fmt.Println(counter()) // 2

	var privateStr string

	hello := func(name string) {
		// fmt.Println("Hello", name+"!")
		privateStr = name
	}

	hello("Hannah")

	privateStr = "zero"
	fmt.Println(privateStr)
	passInStringInput(hello, "Michaela")

	jacuzzi := func(thing string) {
		revised := make([]byte, len(thing))
		for i := len(thing) - 1; i >= 0; i-- {
			revised[len(thing)-i-1] = thing[i]
		}

		fmt.Println(string(revised))
	}
	passInStringInput(jacuzzi, "hello")
	// privateStr = "zero"
	fmt.Println(privateStr)
}

func createCounter() func() int { // if you can output functions, can you pass in functions???
	i := 0
	return func() int {
		i++
		return i
	}
}

func passInStringInput(function func(h string), function_input string) {
	execute := function
	execute(function_input)
}
