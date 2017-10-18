package main

// https://discuss.leetcode.com/topic/30058/a-simple-solution-in-python-based-on-dfa

// 'b': blank
// '.': dot
// 'd': digit (0-9)
// 'e': e
// 's': sign (+/-)
func validNumber(s string) bool {
	state := []map[byte]int{
		// index 0
		map[byte]int{}, // 0
		// state 1, initial state (loop if blank)
		map[byte]int{'b': 1, 's': 2, 'd': 3, '.': 4},
		// state 2, found sign (expect digit/dot)
		map[byte]int{'d': 3, '.': 4},
		// state 3, found digit (loop until non-digit)
		map[byte]int{'d': 3, '.': 5, 'e': 6, 'b': 9},
		// state 4, found dot (expect only digit)
		map[byte]int{'d': 5},
		// state 5, after dot
		map[byte]int{'d': 5, 'e': 6, 'b': 9},
		// state 6, found e (export sign/digit)
		map[byte]int{'s': 7, 'd': 8},
		// state 7, found s after e (expect only digit)
		map[byte]int{'d': 8},
		// state 8, // loop digit
		map[byte]int{'d': 8, 'b': 9},
		// state 9, // loop blank
		map[byte]int{'b': 9},
	}

	curState := 1
	for _, c := range []byte(s) {
		if c >= '0' && c <= '9' {
			c = 'd'
		} else if c == '+' || c == '-' {
			c = 's'
		} else if c == ' ' || c == '\t' || c == '\n' {
			c = 'b'
		}

		if _, exists := state[curState][c]; !exists {
			return false
		}

		curState = state[curState][c]
	}

	// the only valid teminial states are
	// state 3, end on digit
	// state 5, end on digit after dot
	// state 8, end on digit after e
	// state 9, end on whitespace
	for _, v := range []int{3, 5, 8, 9} {
		if curState == v {
			return true
		}
	}

	return false
}
