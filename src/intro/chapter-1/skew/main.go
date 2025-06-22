package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter" // include line graph support and many other graphs
	"gonum.org/v1/plot/vg"      // support for vector graphics and image generation
)

// MinimumSkew
// Input: a (DNA) string genome
// Output: a slice of indices corresponding to the indicies in the genome where the skew array att ains a minimum value
func MinimumSkew(genome string) []int {
	indices := make([]int, 0)

	//make the skew array
	array := SkewArray(genome)

	m := MinIntegerArray(array)

	for i, val := range array {
		if val == m {
			// I found an index :)
			indices = append(indices, i)
		}
	}

	return indices
}

func MinIntegerArray(list []int) int {
	if len(list) == 0 {
		panic("Error: empty list.")
	}

	m := list[0]

	for _, val := range list {
		if val < m {
			m = val
		}
	}

	return m
}

// SkewArray
// Input: A DNA string genome
// Output: A slice of integers corresponding to the "G-C" skew of the genome at each position --> keep track of the difference between the total number of G's and C's in midst of raning the genome!
func SkewArray(genome string) []int {
	n := len(genome)

	array := make([]int, n+1)

	array[0] = 0 // redundant but here just for the sake of thinking / pseudocode

	//range over remaining values and set i-th value of array

	for i := 1; i < n+1; i++ {
		array[i] = array[i-1] + Skew(genome[i-1])
	}

	return array
}

//Skew
//Input: a symbol
//Output: 1 (if symbol is G), -1 (if symbol is C), and 0 otherwise

func Skew(symbol byte) int {
	if symbol == 'G' {
		return 1
	} else if symbol == 'C' {
		return -1
	} else {
		return 0
	}

	//default
	return 0 //biological fact that the circular bacterial genome (plasmic) in 1/2 of the genome before the replication sequence has a depletion of 'G' when compared to 'C' --> after the replication there would be more 'G' than 'C'

	// right now, interested in the position where the skew array goes down and back down, the minimum skew!
}

func main() {
	fmt.Println("The skew array.")

	url := "https://bioinformaticsalgorithms.com/data/realdatasets/Replication/E_coli.txt"

	//grab response from site
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close() // close connection later

	//status OK?
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Received non-OK status %v", resp.Status)
	}

	//access slice of symbols in file
	genomeSymbols, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	//convert genome to string

	genome := string(genomeSymbols)
	fmt.Println("Genome read. It has", len(genome), "total nucleotides")

	EcoliSkewArray := SkewArray(genome)
	minSkewPosition := MinimumSkew(genome)

	firstPosition := minSkewPosition[0]

	fmt.Println("The minimum skew value of", EcoliSkewArray[firstPosition], "occurs at positions", minSkewPosition) // The minimum skew value of -13919 occurs at positions [3923620 3923621 3923622 3923623] --> four consecutive positions of the genome!
	//This is somewhat disappointing --> humans are visual creatures,,, any string has a minimum skew!

	//I am going to make a plot, a Skew Plot
	//--> get pre-written code for this purpose

	//draw the skew diagram
	MakeSkewDiagram(EcoliSkewArray)

	fmt.Println("Skew diagram draw! Exiting normally.")
}

// MakeSkewDiagram
// Input: A skew array
// Output: (none)
// Draws the skew diagram of the skew array to an image and saves to file.
func MakeSkewDiagram(skewArray []int) {
	p := plot.New() // creates a new plotter object

	p.Title.Text = "Skew Diagram"
	p.X.Label.Text = "Genome Position"
	p.Y.Label.Text = "Skew value"

	//remove legend
	p.Legend.Top = false

	// make a collection of points associated with each skew value
	points := make(plotter.XYs, len(skewArray)) // this is n+1 total points bc the skew array has one extra value

	//set the X and Y value of each point
	for i, skewValue := range skewArray {
		points[i].X = float64(i)
		points[i].Y = float64(skewValue)
	}

	//connect the dots!
	line, err := plotter.NewLine(points)

	if err != nil {
		panic(err)
	}

	//add our line to the plot
	p.Add(line)

	//draw to an image

	//first, set a unit of length
	unitOfLength := vg.Centimeter

	//make label fonts bigger
	p.X.Label.TextStyle.Font.Size = 3 * unitOfLength
	p.Y.Label.TextStyle.Font.Size = 3 * unitOfLength
	p.Title.TextStyle.Font.Size = 4 * unitOfLength
	p.X.Tick.Label.Font.Size = 2 * unitOfLength
	p.Y.Tick.Label.Font.Size = 2 * unitOfLength

	//save my plot to a PNG
	err = p.Save(100*unitOfLength, 60*unitOfLength, "skewDiagram.png") // w * h

	if err != nil {
		panic(err)
	}
}
