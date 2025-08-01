package main

import (
	"fmt"
)

func main() {
	// fmt.Println("Genome assembly!")

	// genomeLength := 10000
	// k := 100

	// originalGenome := GenerateRandomGenome(genomeLength)

	// reads := KmerComposition(originalGenome, k)

	// reads = ShuffleStrings(reads)

	// assembledGenome := GreedyAssembler(reads)

	// fmt.Println(assembledGenome)

	// // if assembledGenome == originalGenome {
	// // 	fmt.Println("Success!")
	// // }

	// //check: does the kmer compositions of assembledGenome equal the kmer composition of the original genome
	// kmers := KmerComposition(assembledGenome, k)
	// if StringSliceEquals(reads, kmers) {
	// 	fmt.Println("Sucessfully assembled genome.")
	// }

	// fmt.Println("Program exiting.")

	// fmt.Println(ScoreOverlapAlignment("atcgt", "asecgt", 1.0, 1.0, 3.0))
	SARSOverlapNetworkMinimizerTrim()

	// pattern := "ATCC"

	// adjList := make(map[string][]string)

	// adjList["ATCC"] = []string{"GTCD", "DSDC", "DSCC", "DAJC"}
	// adjList["GTCD"] = []string{"A", "B", "C", "D"}

	// fmt.Println(GetExtendedNeighbors(pattern, adjList, 2))
}

func SARSOverlapNetwork() {
	fmt.Println("Read in the SARS-CoV-2 genome.")

	genome := ReadGenomeFromFASTA("Data/SARS-CoV-2_genome.fasta")

	fmt.Println("Genome read. Sampling reads.")

	//sample some reads
	readLength := 150
	probability := 0.1
	reads := SimulateReadsClean(genome, readLength, probability)

	fmt.Println("Reads generated! Building overlap network.")

	fmt.Println("Now, make the minimizer map.")

	k := 10
	windowLength := 20

	minimizerDictionary := BuildMinimizerMap(reads, k, windowLength)

	fmt.Println("Minimizer map made. It contains", len(minimizerDictionary), "total keys.")

	match := 1.0
	mismatch := 1.0
	gap := 5.0

	threshold := 40.0 //there's also the additional concern that some parts of the same genome will just randomly overlap

	adjList := MakeOverlapNetworkMinimizers(reads, minimizerDictionary, match, mismatch, gap, threshold)

	fmt.Println("Overlap network made!")

	fmt.Println("The network has", len(adjList), "total keys.")

	fmt.Println("The average outdegree is", AverageOutDegree(adjList))

}

func SARSOverlapNetworkMinimizerTrim() {
	fmt.Println("Read in the SARS-CoV-2 genome.")

	genome := ReadGenomeFromFASTA("Data/SARS-CoV-2_genome.fasta")

	fmt.Println("Genome read. Sampling reads.")

	//sample some reads
	readLength := 150
	probability := 0.1
	reads := SimulateReadsClean(genome, readLength, probability)

	fmt.Println("Reads generated! Building overlap network.")

	fmt.Println("Now, make the minimizer map.")

	k := 10
	windowLength := 20

	minimizerDictionary := BuildMinimizerMap(reads, k, windowLength)

	fmt.Println("Minimizer map made. It contains", len(minimizerDictionary), "total keys.")

	match := 1.0
	mismatch := 1.0
	gap := 5.0

	threshold := 40.0 //there's also the additional concern that some parts of the same genome will just randomly overlap

	adjList := MakeOverlapNetworkMinimizers(reads, minimizerDictionary, match, mismatch, gap, threshold)

	fmt.Println("Overlap network made!")

	fmt.Println("The network has", len(adjList), "total keys.")

	fmt.Println("The average outdegree is", AverageOutDegree(adjList))

	maxK := 3

	fmt.Println("Now trimming network by removing transitivity with a maxK value of", maxK)

	trimmedAdjacencyList := TrimNetwork(adjList, maxK)

	fmt.Println("Graph has been trimmed")

	fmt.Println("Graph", len(trimmedAdjacencyList), "total reads")

	fmt.Println("Average outdegree is", AverageOutDegree(trimmedAdjacencyList))
}

/*
How would you select a good threshold? Would it change if the reads were longer? If the error rate were higher?

Higher error rate = lower the threshold
Longer string = raise the threshold

The error rate would probably increase or even if not, cause a lot more errors once the size of the genome increases
This would require that the threshold be turned down a lot more

More penalty to indel rate = more mismatches the threshold should be lowered



*/

/*

We are looking for maximal non-branching paths which are one in one out (our contigs!)

HELP
 ELPN
   PNOT
     OTJU
       JUST
         STAN
	       ANYB
	         YBOD
	           ODYH
	             YHEL
		          HELP

*/
