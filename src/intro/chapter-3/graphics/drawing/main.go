package main

import (
	"fmt"

	"github.com/michaelx0281/Computational-Biology-/src/intro/chapter-3/graphics/canvas"
)

func main() {
	fmt.Println("Drawing a head.")

	DrawFace()
}

/*
Understanding more about canvs.go, which is our engine for drawing.

Go provides a comprehensive library for drawing via the draw2d directory that was installed as part of the setup for this code along.

This directory is not the easiest exactly for a beginner to use direclty, which is why there was some code provided in the canvas folder.

//Special thanks to Car Kingsford and Hannah Kim for putting together the canvas package!//

The cavas p ackage allows us to conceptualize the 'canvas' object. The canvas can be be through of as a rectangular window with a specified width and height.

This is implemented via OOP.

The width and height of the canvas are measured in p ixels, where a single pixel is a small point in the image which will be colored with a single color.

In the RGB color mdoel, every rectangular pixel on a computer screen emits a single color formed as a mistrue of differeing amounts of the three primary colors of light: red, green, and blue. THe intensity determined by the normalized scaling of 0-255

There are a few colors which are shown in the figure below along with the RGB equivalents.const

There are present canvas funciotns which we would need in the next code along (and which are listed below!)

- CreateNewCanvas()
- MakeColor()
- SetFillColor()
- Clear()
- Fill()
- ClearRect()
- Circle()
- GetImage()

*/

// DrawFace takes no inputs and returns nothing.
// It uses the canvas package to draw a face to a file
// in the output folder.
func DrawFace() {
	//create the canvs
	c := canvas.CreateNewCanvas(1000, 2000)

	// fill the canvas as black
	black := canvas.MakeColor(0, 0, 0)
	c.SetFillColor(black)
	c.Clear() // now I wonder what this does! Lets have a mini excursion sometime soon!

	// this has created a "thrilling" black rectangle! Amazing!

	//the standard here uses the top-left as the origin. X coordinates are still relatively normal, however, the Y coordinates are flipped as moving downwards is seen as high positive coefficients!

	//we will be drawing the head as a white circle in the next step!

	white := canvas.MakeColor(255, 255, 255)
	c.SetFillColor(white) // this doesn't fill until fill is called!

	//Circle() takes in the x- y- coords as well as the radius of the circle that is being produced! (float64)

	//The head should ike in the toip part of the rectangle, halfway across from lef tot right. Because the canvas is 1000 px wide, we kow that the x-coord of the circle would be 500.

	//In order to have the same amount of space on the top, left, and right, let's make its y-coord 500, as well. The radius will be set to 200

	c.Circle(500, 500, 200)
	c.Fill() // I am very curiuous as to how it knows what to fill, the inside of the shape or the outside!

	//Adding facial features to the head

	c.SetFillColor(black)
	c.Circle(500, 550, 10)
	c.Fill()

	c.Circle(425, 475, 15)
	c.Circle(575, 475, 15)
	c.Fill()

	red := canvas.MakeColor(255, 0, 0)
	c.SetFillColor(red)
	c.ClearRect(400, 600, 600, 620)

	// save the image associated with the canvas
	c.SaveToPNG("fun.png")
}

//this is about as far as we go for the introduction of filling a canvas!!

//start to look into the actual libraries yourself and perhaps create you own functions too!

//For an interesting code challenge / project! --> utilize this software to create the arrow representations of each of the different alignment mappings!

//Also, try to think more about genomic assembly by tmr
