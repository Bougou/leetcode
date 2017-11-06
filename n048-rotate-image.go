package main

func rotate(matrix [][]int)  {
	// reverse up to down
	l := len(matrix)
	for i := 0; i < l/2; i++ {
			tmp := matrix[i]
			matrix[i] = matrix[l-i-1]
			matrix[l-i-1] = tmp
	}
	
	for i := 0; i < l; i++ {
			// j start from i
			for j:= i; j < l; j++ {
					tmp := matrix[i][j]
					matrix[i][j] = matrix[j][i]
					matrix[j][i] = tmp
			}
	}
}