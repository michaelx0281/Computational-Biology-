package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to hierarchy demo!")
	Parent("hello_world/one")

	ParentChild("example_one", "child_two", "child_three")

	Write("names/michaela.txt", []byte("ex_1:val_1, ex_2:val_2"))

}

// Step One: Make a singular parent folder
func Parent(name string) {
	os.Mkdir(name, 0o755)
} //okay! This works

func ParentChild(name string, children ...string) {
	for _, child := range children {
		os.MkdirAll(name+"/"+child, 0o755)
	}
} //okay! this works too! Beautiful!!

func Write(path string, data []byte) { //can only write to existing folders and etc
	os.WriteFile(path, data, 0o755)
}

//This is relatively simple to use...make sure to adhere to this style in the future as well for all of the data/record files

//as for creating the respective Go types for each of the data info that I would like to collect, that seems like a little bit of a pain...but lets do that soon as well

//I might want to consider creating status messages / codes as well --> but that is mostly for the future
