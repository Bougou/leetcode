package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	rowLen := len(matrix)
	if rowLen == 0 {
		return false
	}
	colLen := len(matrix[0])

	fmt.Println(colLen)
	fmt.Println(rowLen)

	low := 0
	high := colLen*rowLen - 1

	for low <= high {
		mid := (low + high) / 2
		r := mid / colLen
		c := mid % colLen
		if matrix[r][c] == target {
			return true
		} else if matrix[r][c] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return false
}
