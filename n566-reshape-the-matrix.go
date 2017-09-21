package main

func matrixReshape(nums [][]int, r int, c int) [][]int {
	origRowLen := len(nums)
	origColLen := len(nums[0])

	if r*c > origRowLen*origColLen {
		return nums
	}

	row := make([][]int, r)

	for i := 0; i < r; i++ {
		col := make([]int, c)

		for j := 0; j < c; j++ {
			n := i*c + j

			origRowIndex := n / origColLen
			origColIndex := n % origColLen

			col[j] = nums[origRowIndex][origColIndex]
		}
		row[i] = col
	}

	return row
}
