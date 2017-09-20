package main

import (
	"strings"
)

func judgeCircle(moves string) bool {
	return strings.Count(moves, "R") == strings.Count(moves, "L") && strings.Count(moves, "U") == strings.Count(moves, "D")
}

func judgeCircle2(moves string) bool {
	x, y := 0, 0

	for _, m := range []byte(moves) {
		switch m {
		case 'R':
			x++
		case 'L':
			x--
		case 'U':
			y++
		case 'D':
			y--
		}
	}

	return x == 0 && y == 0
}

func judgeCircle3(moves string) bool {
	var path = map[byte][2]int{
		'R': [2]int{1, 0},
		'L': [2]int{-1, 0},
		'U': [2]int{0, 1},
		'D': [2]int{0, -1},
	}

	x, y := 0, 0
	for _, m := range []byte(moves) {
		x += path[m][0]
		y += path[m][1]
	}

	return x == 0 && y == 0
}
