package main

func createMatrix(rows, cols int) [][]int {
	// ?
	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols)
		for j := range matrix[i] {
			matrix[i][j] = i * j
		}
	}
	return matrix
}

