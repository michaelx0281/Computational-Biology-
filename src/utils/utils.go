package utils

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

func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}

//Challenge, lets make this generalizable to any array!

func RemoveDuplicatesFromArray[T comparable](arr []T) []T { //new thing that I learned today about Go, if T is within parameters, inference is enabled. Otherwise, use the []() syntax
	//will delete indices that repeat after every first occurence

	//First, create a map
	occurenceMap := make(map[T]bool) //default value is false

	for index := len(arr) - 1; index >= 0; index-- { //remember, when deleting slices like this, this shifts all indices to the left. If you keep moving right ward, you will move out of bounds. Subtract and move leftwards instead!
		element := arr[index]
		//cut out the current index if it's been marked already!
		if occurenceMap[element] == true {
			arr = append(arr[:index], arr[index+1:]...) //I usually use the '+' operator, but maybe this would work better!
		}
		occurenceMap[element] = true
	}

	return arr
}
