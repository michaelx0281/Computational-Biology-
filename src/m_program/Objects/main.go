package main

type Node struct {
	name           string
	age            float64
	child1, child2 *Node
}

type Tree struct {
	nodes []*Node
}

//why does this not create infinite memory issue?
//pointers default value = nil

type Circle struct {
	c_radius float64
	x, y     float64
}

type Rectangle struct {
	length, width float64
	x, y          float64
	rotation360   int
}

func (c Circle) Area() float64 {
	return c.c_radius * c.c_radius * 3.0
}

func (r Rectangle) Area() float64 {
	return r.length * r.width
}

func (c *Circle) Translate(a, b float64) {
	c.x += a
	c.y += b
}

func (r *Rectangle) Translate(a, b float64) {
	r.x += a
	r.y += b
}

// func CreateCircle(radius, x, y float64) Circle {

// }

//Write a circle and reactangle methods Scale() scaling the size of each shape by a factor f
//The area will be scaled by f

func (c *Circle) Scale(f float64) {
	c.c_radius *= f
}

func (r *Rectangle) Scale(f float64) {
	r.length *= f
	r.width *= f
}

// Remember, with slices, they are pass by reference
// slices are pointers to the first index of an array
func ChangeFirst(list []int) {
	list[0] = 42
}

//in both cases, a copy of what is passsed into the function is created

func main() {
	var t Tree

	t.nodes = make([]*Node, 5)

	// var t Tree

	// t.nodes = make([]Node, 3)
	// t.nodes[0].name = "Mother"
	// t.nodes[1].name = "kid1"
	// t.nodes[2].name = "kid2"

	// var b int = -1
	// var a *int //a is a reference to the location of an integer

	// a = &b //now a will tell you where in the ram b is

	// *a *= 2

	// fmt.Println(b)

	// var c Circle
	// c.x = 2.0
	// c.y = -3.7
	// // c = c.Translate(1.0, 1.0)

	// // we can fix our issue with pointers to our objects

	// var pointerToC *Circle

	// //pointerToC has value nil
	// pointerToC = &c

	// (*pointerToC).x = 43.2
	// // this means de reference ... go in and access the object it points to

	// fmt.Println("New center coordinates", c.x, c.y)

	// //pointer dereference is not needed in go
	// pointerToC.x = 3.0

	// pointerToC.Translate(-3.0, 1.7)

	// // Go doesn't care if you use a pointer or not, it just cares that you use the correct notation
	// c.Translate(4.34, 34.42) // Go automatically creates a pointer over here!

	// a := make([]int, 5) // this creates an array of length 5 and creates a pointer to it (actually is a little more than that)

	// ChangeFirst(a) // a is a pointer to an array

	// fmt.Println(a[0])

	// what gets printed? 0 or 42? // 42 because slices are pass by reference and not pass by value

}

// Pointers are often used to represent connections or edges in a network
// pointers may represent a one way edge

/*
func UPOGMA (D, speciesNames) {
	numLeaves <- CountRows(D)
	//count the number of rows (create the cluster array based on this)
	t <- InitializeTree(speciesNames)
	clusters <- InitializeClusters(t)
	for every integer p from numLeaves to 2*numLeaves-2 (n-1 internal nodes theorem)
		row, col, value <- FindMinElt(D)
		t.nodes[p].age <- val/2
		t.nodes[p].child1 <- clusters[row]
		t.nodes[p].child2 <- clusters[col]
		clusterSize1 <- CountLeaves(t.nodes[p].child1) // reference to some row --> counts the number of leaves beneath --> use Recursion!
		clusterSize2 <- CountLeaves(t.nodes[p].child2)

		D <- AddRowCol(D, clusterSize1, clusterSize2, row, col)
		D <- DelRowCol(D, row, col)

}

*/

/*

How to implement a function that counts the number of leaves?

You would first start with the root node
{

age,
name,
child1, child2 *Node

counter update for each Node, and recurse until child 1 AND child 2 is nil

func func(Node) int //counting LEAVES only

if n.Child1 == nil && n.Child2 == nil {
	return 1
} else if n.Child1 == nil {
	return func(n.Child2)
} else if n.Child 2 == nil {
	return func(n.Child1)
} else {
	return func(n.Child1) + func(n.Child2)
}


//this is a good recursive function because no two things are calling count leaves of the same node
*/
