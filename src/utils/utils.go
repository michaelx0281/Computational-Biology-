package utils

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func PrintIntListSpaceSeparated(data []int) {
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

func HandleError(e error, msg string) {
	if e != nil {
		panic("Encountered error:" + e.Error() + "." + " " + msg)
	}
}

func HandleErrorLog(e error, msg ...string) {
	if e != nil {
		if len(msg) == 1 {
			log.Fatal("Encountered error %s", msg[0])
		}

		if len(msg) == 2 {
			log.Fatal("Encountered error %s could not %s", msg[0], msg[1])
		}

		if len(msg) >= 3 {
			panic("Too many error msgs. Re-check code implementation.")
		}
	}
}

func HandleFatalFileCreationError(e error, name, msg string) {
	if e != nil {
		log.Fatalf("Encountered fatal error, could not write %v to file %s", e, name)
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

// Challenge: new generic!

func NumberUniqueElements[T comparable](arr []T) int {
	occuranceTable := make(map[T]bool)
	count := 0
	for _, value := range arr {
		if occuranceTable[value] == true {
			continue
		}
		count++
		occuranceTable[value] = true
	}

	return count
}

// Challenge: write to file

func WriteBytesToTxTFile(name string, data []byte) {
	//if you want to write to a folder, still need os.MkdirAll and os.Join in order to do that
	fmt.Println("Length of data is", len(data))
	file, err := os.Create(name)

	HandleFatalFileCreationError(err, name, "System was not able to create file. Check the input name string for formatting.")

	data = []byte(AddSpacesToString(string(data))) // don't let this fool you, this is actually a type cast and nothing fancy lol

	bytesWritten, err2 := file.Write(data)

	HandleErrorLog(err2, "writing to file failure, ", "write file to"+name)

	fmt.Printf("%d bytes written to file %s", bytesWritten, name)

}

func WriteIntsToTxTFile(name string, data []int) {
	//if you want to write to a folder, still need os.MkdirAll and os.Join in order to do that
	fmt.Println("Length of data is", len(data))
	fmt.Println("hi")
	file, err := os.Create(name)

	fmt.Println("hi2")
	HandleFatalFileCreationError(err, name, "System was not able to create file. Check the input name string for formatting.")

	defer file.Close() //maybe there were errors bc the file tried to close too soon?

	fmt.Println("List before null byte list.", IntListToString(data))

	fmt.Println("List which produced null byte error.", AddSpacesToString(IntListToString(data)))
	fmt.Println("waoifejwoifjp")

	data = StringToIntList(AddSpacesToString(IntListToString(data)))
	bytesWritten, err3 := file.Write(IntListToByteList(data))

	HandleErrorLog(err3, "writing to file failure, ", "write file to"+name)

	fmt.Printf("%d bytes written to file %s", bytesWritten, name)
}

func IntListToString(list []int) string { //run a different protocal that has an array of strings in this format str[n][2] and return errors otherwise
	// making each index-element type casted into a byte does NOT work!
	s := make([]byte, len(list)) //seems like I would have to make a list of strings instead

	// 1 => strconv to a ==> cast as byte = 49
	// -1 => strcov to a ==> cast as byte = 45 49 (this is because there are two characters in '-1')

	for i, element := range list {

		stringedElement := strconv.Itoa(element)
		fmt.Println(stringedElement)
		// fmt.Println(element)
		fmt.Println("Element:", stringedElement)
		fmt.Println("The same thing but with 1", []byte(strconv.Itoa(1)))
		fmt.Println("-", []byte(strconv.Itoa(-1)))
		slice := []byte(stringedElement)

		//check if one character
		if len(slice) == 1 {
			s[i] = slice[0]
			continue
		}
		fmt.Println(slice[0])
		fmt.Println(slice[1])

		panic("Length element in list is not 1. String to byte conversion failed. Length of slice is " + strconv.Itoa((len(slice))))
	}
	return string(s)
}

func IntListToByteList(list []int) []byte {
	s := make([]byte, len(list))
	for i, element := range list {
		s[i] = byte(element)
	}
	return s
}

func StringToIntList(s string) []int {
	list := make([]int, len(s))
	for i, char := range s {
		stringedChar, err := strconv.Atoi(string(char))

		HandleError(err, "Encountered error could not convert string-cast character to int at index "+strconv.Itoa(i))
		list[i] = stringedChar
	}

	return list
}

func AddSpacesToString(str string) string {
	newStr := make([]byte, len(str)*2-1)

	for i := range newStr {
		//this is the end of newStr
		if i == len(newStr)-1 {
			newStr[i] = str[len(str)-1]
			continue
		}
		//this is the start of newStr
		if i == 0 {
			newStr[i] = str[i]
			continue
		}

		if i%2 == 0 {
			newStr[i] = str[i/2] //this works better now!
			continue
		} else {
			newStr[i] = ' '
		}
	}

	return string(newStr)
}
