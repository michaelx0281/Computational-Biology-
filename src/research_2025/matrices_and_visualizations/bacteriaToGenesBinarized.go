package main

type Axis []string

/*
How should I plan this out?

There are two axes: Species (+variants) and Genes

You want to check if every species is able to be matched to a gene.

You need to be able to pull from a database of information that tells you whether the gene was found under the specific species

You need to use MyGenes to check! --> never mind, I'm moving away from using this as I don't particularly like this software that much.const

Watching some YT Tutorials on utilizing the API on NCBI for this exact purpose right now. Specifically, I amn looking into the E-utilities packages using estrez (most likely and hopefully)

recording this here in case I would need it for later:

eutils.ncbi.nlm.nih.gov --> there is a lot of documentation and other stuff here

Text Search --> databases such as PubMed, Nucleotide, Gene
Downlaoding records in various formatts
Linking between records from different databases

There are 7 cgis using the same base url

baseURL/esearch.fcgi?db=nuccore&term=mouse[orgn]
eSearch -> accepts text query in a specified database and returns integer identifiers (UIDs) for any matching records
--> this would return UIDs for all mouse seqs in the nucleotide database

Aside from eFetch, all of the utilities return data in the XML formatting

<Count> --> number of total records fetched
<RetMax> max you want to pull up at the moment
<RetStart> the index that you want it to start pulling from (starts at 0)

eSummary

baseURL/esummary.fcgi?db=nuccore&id=49619226,49615287
--> input: UID   output: document summaries in XML

eFetch

baseURL/efetch.fcgi?db=nuccore&id=49619226,49615287
--> like eSummary, except it returns full records in XML

&retmode i.e. Text (what other options are there? this is kind of interested but also pretty confusing)
&rettype --> the record format you want it in, i.e. FASTA (lets use this one)

not all dbs are supported by eFetch, look further are Table 1, Chp 4 for the documentation to make further sense of everything within

alternative to dealing with large lists of UIDs

two parameters: QueryKey, WebEnv


Usage Guidelines:

make no more than 3 req per second

Run large jobs on weekends or between 5 pm and 9 AM EST

include &tool and &email in all requests

--> this would allow ncbi to contact before blocking access (if excessive)

register &tool and &email values, ask questions, by writing to NCBI at

eutilities@ncbi.nlm.nih.gov

*/

func BacteriaToGenesBinarized(x Axis, y Axis) map[string]bool {

	/*
		Notes of implementation.

		TODO: change map[string]bool to something closer to map[string, string]bool ==> probably using a struct(string, string) as fields

		One of the goals: make as few queries as possible.

		1) Generate list of all necessary UIDs of each organism

		--> The UIDs for each probiotic of interest would be different across different databases

		Try: UIDs of listed probiotics from research paper on research2025 google doc --> ESearch on Taxonomy with usehistory=y (grab the key and webenv) -->
			--> ELink (would prefer [Category][taxon][gene] in terms of hierarchy for writing (figure it out!))
				(
					Priority:
						Gene
					Sidebar:
						Nuccore(dna, rna, mrna)
						Protein(proteomics)

				)
			--> make function that:
					Generate list of top 20-50 most popular genes (extend as needed with the size of the dataset)
					--> BONUS: visualize

					PRIORITY:
					Make matrix of probiotics vs genes
					--> multiple matches associated to better colors, I suppose

					make a bucket list of colors

					can try to use assembly or some other annotated gene database to make further generalizations of the specificity of usage in relation to immune --> other benefits

					(might need to learn some R visualizaiton techniques for this)

					Make matrix of probiotics vs extended list of genes (retrieve using Assembly-->SIDEBAR)

			--> narrow down to 5-10% list in terms of strains / organisms of interest (do the 2-5 first and use clustalomega to perform quality control tests and whatnot, then build off of it) --> at this point, look back at the document and try to follow the instructions strictly

		Probio and Probio-ichnos --> look at these databases which are specifically curated towards probiotics


	*/

	return make(map[string]bool)
}

/*

I should spend the time to work on this later

*/
