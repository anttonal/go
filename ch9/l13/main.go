package main

func createMatrix(rows, cols int) [][]int {
	matrix := [][]int{}

	if rows != 0 && cols != 0 {
		for i := range rows {
			tempCol := []int{}
			for j := range cols {
				tempCol = append(tempCol, i*j)
			}
			matrix = append(matrix, tempCol)
		}
	}
	return matrix
}
