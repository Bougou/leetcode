package main

func lengthOfLastWord(s string) int {
	e := -1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			e = i
			break
		}
	}

	if e == -1 {
		return 0
	}

	b := e
	for i := b - 1; i >= 0; i-- {
		if s[i] != ' ' {
			b = i
		} else {
			break
		}
	}

	return e - b + 1
}

func lengthOfLastWord2(s string) int {
	res := 0 // store the last word

	cur := 0 // store len of current word
	for _, c := range s {
		if c != ' ' {
			cur += 1
			res = cur
		} else {
			cur = 0
		}
	}

	return res
}
