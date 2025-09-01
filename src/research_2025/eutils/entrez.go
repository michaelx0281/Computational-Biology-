package eutils

/*
	This is something that I am going to use to further automate the process of retrieval of records from different databases based on relevant search terms.
*/

type fcgi struct {
	base_url string
	db       Database
	function Function
	term     Term
	tool     Tool
	email    Email
}

type Email string
type Tool string

const (
	personal Email = "&email=sonic0281mx@gmail.com"
)

const (
	tool Tool = "&tool=michaelaistired"
)

func (f fcgi) Assemble() string {
	url := append([]byte(f.base_url), []byte(f.function.String())...)
	url = append(url, []byte(f.db.String())...)
	url = append(url, []byte(f.term.URLSyntax())...) //maybe i should have a separate group of functions that returns in []byte instead to make it easier
	url = append(url, []byte(string(f.tool))...)
	url = append(url, []byte(string(f.email))...)

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
		"db=protein",
		"db=nuccore",
		"db=pubmed",
		"db=pmc",
		"db=gene",
		"db=nlmcatalog",
		"db=omia",
		"db=structure",
		"db=taxonomy",
		"db=sra",
		"db=biosample",
		"db=bioproject"}[d-1]

	// I am not sure as to how exactly this works at the moment, but I will be happy as long as it works
}

/* Functions */
type Function int

const (
	ESearch Function = iota + 1
	ESummary
)

func (f Function) String() string {
	return [...]string{
		"esearch.fcgi?",
		"esummary.fcgi?"}[f-1]
}

/* Terms */

type Term string

// 5%B = [
// 5%D = ]

/* The above is necessary when you need to separate vastly different types of terms */

func (t Term) URLSyntax() string {
	return "&term=" + SpliceInsert(string(t))
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

func Fcgi(f Function, db Database, t Term) fcgi {

	entrez := fcgi{base_url: "http://eutils.ncbi.nlm.nih.gov/entrez/eutils/", db: db, function: f, term: t, tool: tool, email: personal}

	return entrez
}
