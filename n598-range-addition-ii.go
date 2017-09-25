package main

func maxCount(m int, n int, ops [][]int) int {
	minX := m
	minY := n

	for _, op := range ops {
		if op[0] < minX {
			minX = op[0]
		}

		if op[1] < minY {
			minY = op[1]
		}
	}

	return minX * minY
}
