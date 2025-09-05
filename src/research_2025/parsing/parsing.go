package parsing

import (
	"encoding/xml"
	"net/http"

	"github.com/michaelx0281/Computational-Biology/src/utils"
)

/*
	For parsing into structs... basic structure looks to be

	type Name struct {
		FieldA string `xml: "XMLField"` --> keep building on it where Name2 / Name3 uses Name
		FieldB string
	}
*/

type Clump struct {
	Clump []Gene `xml:"Entrezgene-Set"` //I think that this makes the most sense at the moment, still unsure as to if this would work, however
}

type Gene struct {
	Gene GeneWrapperPlural `xml:"Entrezgene_gene"` //change this to something else if this doesn't seem to work
}

type GeneWrapper struct { //expand to synonyms too
	Name string `xml:"Entrezgene_gene>Gene-ref"`
	UID  string `xml:"Gene-ref_db>Dbtag>Object-id>Object-id_id"`
}

type GeneWrapperPlural struct {
	Name      string   `xml:"Gene-ref_locus"`
	UID       string   `xml:"Gene-ref_db>Dbtag>Object-id>Object-id_id"`
	Neighbors []string `xml:"Gene-ref_syn"`
}

/* This is parsing for a specific gene... there are additional structures for queries with multiple uids */
func ParseGeneRecord(url string) map[string]string {
	resp, err := http.Get(url)

	utils.HandleErrorLog(err, "HTTP req failed.") //need to improve on this section next time

	defer resp.Body.Close()

	var output Clump
	err2 := xml.NewDecoder(resp.Body).Decode(&output) //storing output from set of genes here at this memory address
	utils.HandleErrorLog(err2, "Failure to parse into Go structs.")

	GeneTable := make(map[string]string)

	for _, g := range output.Clump {
		GeneTable[g.Gene.Name] = g.Gene.UID

	}

	//should i make a table based on if genes are neigbors with each other?

	return GeneTable
}

/*
	The intent of the below function is to provide a map[uid]term so that the UIDs could later be fed to ESummary or EFetch.

	It may be better to use the query key and env in some cases; however, I still need a function that would properly map each key to something more of significance to me once graphed
*/

//next step for these, would be to write the `xml:xxx` tags for each so that the decoder knows which fields to associate with which Go struct

type Query struct {
	list []Identifier
}

type Identifier struct {
	id   UID
	name taxon
}

type taxon string
type UID int

func ParseSearchUIDToTerm(url string) map[int]string {
	resp, err := http.Get(url)

	utils.HandleErrorLog(err, "Http req. failed") //now with multiple of them, I needt to be able to better identify which function failed. I can include a status code and have a function that could interpret status codes in order to return relevant infomration about what specifically failed --> but this would probablywait for later. Project for after all of the essentials are done.

	defer resp.Body.Close()

	//define a pointer to a storage object, and feed the memory address to the decoder

	var q Query
	err2 := xml.NewDecoder(resp.Body).Decode(&q)
	utils.HandleErrorLog(err2, "Trouhle with decoding...")

	id := make(map[int]string)

	for _, orgn := range q.list {
		id[int(orgn.id)] = string(orgn.name)
	}

	return id
}

// func GenerateNeighborhoods() {

// }

// func NeigbhorsTable() {
// 	type GeneComparison struct {

//idk if this is really the most efficient when it is already paired (parent:[]neighbors) --> but I think that this is most likely to capture the fullest picture???

//--> overlapping neighborhoods?

// 		parent_gene   string
// 		neigbors_gene string
// 	}

// 	GeneNeighbors := make(map[GeneComparison]bool)

// 	//TODO: logic for associating two different strings with one another

// 	//later--> look into the average in-degree of each of the genes to determine which ones are most popular base on a relational scale

// 	//watch back on the video about definition of different neighborhoods (the smallest number of arrows in between neighborhoods, most within--> need to further crystalize this definitition)
// }

//func to parse Taxon?
//ELink --> fromdb=taxonomy, db=gene, uid: --> get a web-env from a text query form the first time
