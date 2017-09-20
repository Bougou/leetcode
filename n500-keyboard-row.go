package main

import (
	"strings"
)

func findWords(words []string) []string {
	var (
		row1 = "qwertyuiop"
		row2 = "asdfghjkl"
		row3 = "zxcvbnm"
	)

	var letterMap = make(map[byte]int)
	for _, c := range []byte(row1) {
		letterMap[c] = 1
	}

	for _, c := range []byte(row2) {
		letterMap[c] = 2
	}

	for _, c := range []byte(row3) {
		letterMap[c] = 3
	}

	var result []string

	for _, word := range words {
		flag := true

		lword := strings.ToLower(word)
		for _, c := range []byte(lword) {
			if letterMap[c] != letterMap[lword[0]] {
				flag = false
				break
			}
		}

		if flag == true {
			result = append(result, word)
		}
	}

	return result
}
