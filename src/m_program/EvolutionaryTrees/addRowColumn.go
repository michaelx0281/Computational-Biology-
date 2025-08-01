package main

//AddRowCol takes a distance matrix Given a DistanceMatrix, a slice of current clusters,
//and a row/col index (NOTE: col > row).
//It returns the matrix corresponding to "gluing" together clusters[row] and clusters[col]
//forming a new row/col of the matrix for the new cluster, computing
//distances to other elements of the matrix weighted according to the sizes
//of clusters[row] and clusters[col].
func AddRowCol(row, col, clusterSize1, clusterSize2 int, mtx DistanceMatrix) DistanceMatrix {
	return mtx
}
