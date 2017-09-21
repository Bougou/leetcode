package main

func islandPerimeter(grid [][]int) int {
	rowLen := len(grid)
	colLen := len(grid[0])

	var (
		numLan   = 0
		numNeigh = 0
	)

	for r := 0; r < rowLen; r++ {
		for c := 0; c < colLen; c++ {
			if grid[r][c] == 1 {
				numLan++

				if (r+1 < rowLen) && grid[r+1][c] == 1 {
					numNeigh++
				}

				if (c+1 < colLen) && grid[r][c+1] == 1 {
					numNeigh++
				}
			}
		}
	}

	return numLan*4 - numNeigh*2
}
