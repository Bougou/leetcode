package main

func exist(board [][]byte, word string) bool {
	if word == "" {
		return true
	}

	for x := 0; x < len(board); x++ {
		for y := 0; y < len(board[0]); y++ {
			// board, word, current x-index, current y-index, current letter index in word
			if exists(board, word, x, y, 0) {
				return true
			}
		}
	}

	return false
}

func exists(board [][]byte, word string, x, y, i int) bool {
	if i == len(word) {
		return true
	}

	if x < 0 || y < 0 || x == len(board) || y == len(board[0]) || word[i] != board[x][y] {
		return false
	}

	board[x][y] = '*'
	e := exists(board, word, x+1, y, i+1) ||
		exists(board, word, x-1, y, i+1) ||
		exists(board, word, x, y+1, i+1) ||
		exists(board, word, x, y-1, i+1)
	board[x][y] = word[i]

	return e
}
