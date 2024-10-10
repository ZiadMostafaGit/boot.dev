package main

import "fmt"

func createMatrix(rows, cols int) [][]int {

	// Create a slice of slices with `rows` number of rows
	matrix2d := make([][]int, rows)

	// Initialize each inner slice to have `cols` columns
	for i := 0; i < rows; i++ {
		matrix2d[i] = make([]int, cols) // Allocate space for the columns in each row
		for j := 0; j < cols; j++ {
			matrix2d[i][j] = j * i // Assign values
		}
	}

	return matrix2d
}

func main6() {

	res := createMatrix(10, 10)

	// Print the matrix
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Printf("%v ", res[i][j])
		}
		fmt.Println()
	}
}
