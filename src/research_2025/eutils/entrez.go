package eutils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/michaelx0281/Computational-Biology/src/utils"
)

/*
	This is something that I am going to use to further automate the process of retrieval of records from different databases based on relevant search terms.
*/

type fcgi struct {
	base_url string
	db       Database
	db_from  Database
	function Function
	term     Term
	tool     Tool
	email    Email
	api_key  Key
}

type Key string

func (k Key) String() string {
	return "&api_key=" + string(k)
}

type Email string
type Tool string

const (
	personal Email = "&email=sonic0281mx@gmail.com"
)

const (
	tool Tool = "&tool=michaelaistired"
)

func (f fcgi) Assemble() string { //if any of the conditionals grow bigger, start to use switch-case statements instead
	if f.function != ELink {
		return withoutELink(f)
	} else {
		return withElink(f)
	}
}

func withoutELink(f fcgi) string {
	url := append([]byte(f.base_url), []byte(f.function.String())...)
	url = append(url, []byte("db="+f.db.String())...)
	url = append(url, []byte(f.term.URLSyntax())...) //maybe i should have a separate group of functions that returns in []byte instead to make it easier
	url = append(url, []byte(string(f.tool))...)
	url = append(url, []byte(string(f.email))...) //appending so much is probably not very efficient, should look into a better way of doing this some other time
	url = append(url, []byte(string(f.api_key.String()))...)

	//other than the repetitiveness of everything, even if it seems like it is a little bit unecessary, I quite like the way that this is structured overall!
	return string(url)
}

func withElink(f fcgi) string {
	url := append([]byte(f.base_url), []byte(f.function.String())...)
	url = append(url, []byte("db="+f.db.String())...)
	url = append(url, []byte("&dbfrom="+f.db_from.String())...)
	url = append(url, []byte(f.term.URLSyntax())...) //maybe i should have a separate group of functions that returns in []byte instead to make it easier
	url = append(url, []byte(string(f.tool))...)
	url = append(url, []byte(string(f.email))...) //appending so much is probably not very efficient, should look into a better way of doing this some other time
	url = append(url, []byte(string(f.api_key.String()))...)

	//other than the repetitiveness of everything, even if it seems like it is a little bit unecessary, I quite like the way that this is structured overall!
	return string(url)
}

/* Database */
type Database int

const (
	Protein Database = iota + 1
	Nucleotide
	PubMed
	PubMed_Central
	Gene
	NLM_Catalog
	OMI
	Structure
	Taxonomy
	SRA
	BioSample
	BioProject
)

func (d Database) String() string { //string version of each enum
	return [...]string{
		"protein",
		"nuccore",
		"pubmed",
		"pmc",
		"gene",
		"nlmcatalog",
		"omia",
		"structure",
		"taxonomy",
		"sra",
		"biosample",
		"bioproject"}[d-1]

	// I am not sure as to how exactly this works at the moment, but I will be happy as long as it works
}

/* Functions */
type Function int

const (
	ESearch Function = iota + 1
	ESummary
	EFetch
	EgQuery
	ELink
	EInfo
	EPost //I don't really see my usages extending past this
)

func (f Function) String() string {
	return [...]string{
		"esearch.fcgi?",
		"esummary.fcgi?",
		"efetch.fcgi?",
		"egquery.fcgi?",
		"elink.fcgi?",
		"einfo.fcgi?",
		"epost.fcgi?"}[f-1]
}

/* Terms */

type Term string

// 5%B = [
// 5%D = ]

/* The above is necessary when you need to separate vastly different types of terms */

func (t Term) URLSyntax() string {
	return "&term=" + SpliceForwardDash(SpliceInsert(string(t))) //considering using a dependency to manage these http encodings...its starting to get a little it annoying
}

func SpliceForwardDash(Text string) string {
	text := []byte(Text)

	n := len(text)

	for i := n - 1; i >= 0; i-- {
		if text[i] == '/' {
			modification := append([]byte{'%', '2', 'F'}, text[i+3:]...)
			text = append(text[:i], modification...)
		}
	}

	return string(text)
}

func SpliceInsert(Text string) string { //might need to add percent encoding, but nothing should break as long as I don't use invalid characters (aka non-alpha )
	text := []byte(Text)

	for i, char := range text {
		if char == ' ' {
			text[i] = '+'
		}
	}

	return string(text)
}

/* Modifiers */ //will not be adding this until much later!

func Fcgi(f Function, t Term, db ...Database) fcgi {

	if len(db) > 2 {
		log.Fatal("Too many databases. 2 MAX limit.")
	} //this error should technically be split between the two scenarios; however, I am too tired to do that rn

	err := godotenv.Load("../../.env") //I had not initially known that this was necessary somehow..
	if err != nil {
		utils.HandleErrorLog(err, "Error loading .env file")
	}

	var entrez fcgi

	if f == ELink {
		entrez = fcgi{
			base_url: "http://eutils.ncbi.nlm.nih.gov/entrez/eutils/",
			db:       db[1],
			db_from:  db[0],
			function: f,
			term:     t,
			tool:     tool,
			email:    personal,
			api_key:  Key(os.Getenv("ENTREZ_API_KEY"))}
	} else {
		entrez = fcgi{
			base_url: "http://eutils.ncbi.nlm.nih.gov/entrez/eutils/",
			db:       db[0],
			function: f,
			term:     t,
			tool:     tool,
			email:    personal,
			api_key:  Key(os.Getenv("ENTREZ_API_KEY"))}
	}

	return entrez
}
