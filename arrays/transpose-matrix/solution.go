package main

func TransposeMatrix(matrix [][]int) [][]int {
  

    // Create a new matrix with dimensions swapped (transposed) and populate it.
    transposedMatrix := make([][]int,len(matrix[0]))
    
    for col:=0;col < len(matrix[0]) ; col++ {
        transposedMatrix[col] = make([]int, len(matrix))
        for row :=0;row <len(matrix); row++ {
            transposedMatrix[col][row] = matrix[row][col]
        }
    }

    return transposedMatrix
}
