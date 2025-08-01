package main

// BinarizeMatrix takes a matrix of values and a threshold.
// It binarizes the matrix according to the threshold.
// If entries across main diagonal are both above threshold, only retain the bigger one.
func BinarizeMatrix(mtx [][]float64, threshold float64) [][]int {
	matrix := Make2D_2[int](len(mtx), len(mtx[0]))

	for i := 0; i < len(mtx); i++ {
		for j := 0; j < len(mtx[0]); j++ {
			if mtx[i][j] >= threshold {

				matrix[i][j] = 1
			} else {
				matrix[i][j] = 0
			}

		}

	}
	for i := 0; i < len(mtx); i++ {
		for j := 0; j < len(mtx[0]); j++ {
			if matrix[i][j] == 1 {
				if (mtx[i][j] >= threshold && (mtx[i][j] > mtx[j][i])) || (mtx[i][j] > threshold && mtx[i][j] == mtx[j][i] && i < j) {
					/**
					if i == 2 && j == 0 {

						fmt.Println(mtx[i][j] >= threshold && mtx[i][j] > mtx[j][i])
						fmt.Println(mtx[i][j], mtx[j][i])
						fmt.Println(mtx[i][j] > threshold && mtx[i][j] == mtx[j][i] && i < j)

					}**/
					matrix[i][j] = 1
				} else {
					matrix[i][j] = 0
				}
			}
		}
	}
	return matrix
}
