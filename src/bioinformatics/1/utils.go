package main

import (
	"fmt"
	"strconv"
)

func PrintListSpaceSeparated(data []int) {
	for i, index := range data {
		if i == len(data)-1 { //at the last step, we do not want a space.
			fmt.Print(strconv.Itoa(index)) //need to use strcov Itoa here!
			continue
		}
		fmt.Print(strconv.Itoa(index) + " ") //type match the int and the spacing
	}
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}
