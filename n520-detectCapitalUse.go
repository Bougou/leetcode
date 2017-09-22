package main

func detectCapitalUse(word string) bool {
	data := []rune(word)

	capCount := 0
	for _, c := range data {
		if 'Z'-c >= 0 {
			capCount++
		}
	}

	if capCount == len(data) ||
		capCount == 0 ||
		(capCount == 1 && 'Z'-data[0] >= 0) {
		return true
	}

	return false
}
