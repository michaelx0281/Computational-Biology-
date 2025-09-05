package main

import (
	"fmt"

	. "github.com/michaelx0281/Computational-Biology/src/research_2025/eutils" //dealing with submodules is such a pain...look to do it later if you want..the github submodule already took me substantial time, don't want to do the same for go modules any honestly
)

func main() {
	// fmt.Println(
	// 	Fcgi(
	// 		ELink,
	// 		Term(CommaSeparatedString(
	// 			"Bacillus mesentericus",
	// 			"Bacillus subtilis",
	// 			"Bacillus velezensis",
	// 			"Bifidobacterium animalis DN-173 010",
	// 			"Bifidobacterium animalis",
	// 			"Bifidobacterium animalis NumRes252/-253",
	// 			"Bifidobacterium animalis subsp. Lactis",
	// 			"Bifidobacterium bifidum",
	// 			"Bifidobacterium breve",
	// 			"Bifidobacterium breve IPLA 20004",
	// 			"Bifidobacterium breve M-16V",
	// 			"Bifidobacterium breve AH1205",
	// 			"Bifidobacterium breve UCC2003",
	// 			"Bifidobacterium bifidum MIMBb23sg",
	// 			"Bifidobacterium bifidum LMG13195",
	// 			"Bifidobacterium lactis HN019",
	// 			"Bifidobacterium longum",
	// 			"Bifidobacterium longum",
	// 			"Bifidobacterium longum AH1206",
	// 			"Bifidobacterium infantis NumRes251",
	// 			"Bifdobacterium infantis",
	// 			"Clostridium butyricum",
	// 			"Escherichia coli",
	// 		)),
	// 		Taxonomy,
	// 		Gene).Assemble())

	fmt.Println(SearchTaxonUIDExample(
		"Bacillus mesentericus",
		"Bacillus subtilis",
		"Bacillus velezensis",
		"Bifidobacterium animalis DN-173 010",
		"Bifidobacterium animalis",
		"Bifidobacterium animalis NumRes252/-253",
		"Bifidobacterium animalis subsp. Lactis",
		"Bifidobacterium bifidum",
		"Bifidobacterium breve",
		"Bifidobacterium breve IPLA 20004",
		"Bifidobacterium breve M-16V",
		"Bifidobacterium breve AH1205",
		"Bifidobacterium breve UCC2003",
		"Bifidobacterium bifidum MIMBb23sg",
		"Bifidobacterium bifidum LMG13195",
		"Bifidobacterium lactis HN019",
		"Bifidobacterium longum",
		"Bifidobacterium longum",
		"Bifidobacterium longum AH1206",
		"Bifidobacterium infantis NumRes251",
		"Bifdobacterium infantis",
		"Clostridium butyricum",
		"Escherichia coli",
	))
}

func TaxonToGene(Terms ...string) {
	Fcgi(
		ELink,
		Term(CommaSeparatedString(Terms...)),
		Taxonomy).Assemble()
}

func SearchTaxonUIDExample(Terms ...string) string { //make a parser for Taxo UID ==> create map[taxon]id in this way
	return Fcgi(
		ESearch,
		Term(SpliceOR(Terms...)),
		Taxonomy).Assemble()
}

func StitchedString(strings ...string) string {
	str := make([]byte, 0)

	for _, s := range strings {
		str = append(str, []byte(s)...)
	}

	return string(str)
}

func SpliceOR(strings ...string) string {
	result := make([]byte, 0)

	for i, str := range strings {
		if i == 0 {
			result = append(result, []byte(str)...)
			continue
		}

		str = " OR " + str //prepending the comma
		result = append(result, []byte(str)...)
	}

	return string(result)
}

func CommaSeparatedString(strings ...string) string {
	result := make([]byte, 0)

	for i, str := range strings {
		if i == 0 {
			result = append(result, []byte(str)...)
			continue
		}

		str = "," + str //prepending the comma
		result = append(result, []byte(str)...)
	}

	return string(result)
}
