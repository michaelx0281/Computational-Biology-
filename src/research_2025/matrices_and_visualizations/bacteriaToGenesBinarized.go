package main

type Axis []string

/*
How should I plan this out?

There are two axes: Species (+variants) and Genes

You want to check if every species is able to be matched to a gene.

You need to be able to pull from a database of information that tells you whether the gene was found under the specific species

You need to use MyGenes to check!

*/

func BacteriaToGenesBinarized(x Axis, y Axis) map[string]bool {
	hashTable := make(map[string]bool)

	return hashTable
}

/*

I should spend the time to work on this later

*/
