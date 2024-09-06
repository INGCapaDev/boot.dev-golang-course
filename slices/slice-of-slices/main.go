package main

func createMatrix(rows, cols int) [][]int {

	matrix := make([][]int, rows)

	for r := 0; r < rows; r++ {
		newRow := make([]int, cols)
		for c := 0; c < cols; c++ {
			newRow[c] = c * r
		}
		matrix[r] = newRow
	}

	return matrix
}
