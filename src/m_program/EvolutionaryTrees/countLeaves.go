package main

// CountLeaves takes a non-nil pointer to a Node object and returns
// the number of leaves in the tree rooted at the node. It returns 1 at a leaf.
func CountLeaves(v *Node) int {
	//base case
	if v.Child1 == nil && v.Child2 == nil {
		return 1
	} else if v.Child1 == nil {
		return CountLeaves(v.Child2)
	} else if v.Child2 == nil {
		return CountLeaves(v.Child1)
	} else {
		return CountLeaves(v.Child1) + CountLeaves(v.Child2)
	}
}
