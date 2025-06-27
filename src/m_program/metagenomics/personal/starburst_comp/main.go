package main

import "gonum.org/v1/plot"

// 	"gonum.org/v1/plot"
// 	"gonum.org/v1/plot/vg" //vector graphics
// 	"fmt"
// encountered problem: there is no native function to visualize and create a sunburst chart within go --> maybe I could try to make one myself --> would require learning how visualization actually works

func main() {
	starburst := plot.New() // creates a new plotter object

	starburst.Title.Text = "Starbust k-Mer diagram"
	starburst.X.Label.Text = "n buckets corresponding to the number of letters"
}
