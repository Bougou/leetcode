package main

func reverseString(s string) string {
	data := []rune(s)
	ret := []rune{}

	for i := len(data) - 1; i >= 0; i-- {
		ret = append(ret, data[i])
	}

	return string(ret)
}
