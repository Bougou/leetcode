package main

import "strings"

func reverseWords(s string) string {
	ret := []string{}

	for _, word := range strings.Split(s, " ") {
		data := []rune(word)
		reverse := []rune{}

		for i := len(data) - 1; i >= 0; i-- {
			reverse = append(reverse, data[i])
		}

		ret = append(ret, string(reverse))
	}

	return strings.Join(ret, " ")
}
