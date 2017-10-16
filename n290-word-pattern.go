package main

import "strings"

func wordPattern(pattern string, str string) bool {
	p := []byte(pattern)
	s := strings.Split(str, " ")

	mc := make(map[byte]string)
	ms := map[string]byte{}

	if len(p) != len(s) {
		return false
	}

	for i := 0; i < len(p); i++ {
		_, exists := mc[p[i]]
		if !exists {
			mc[p[i]] = s[i]
		}
		if mc[p[i]] != s[i] {
			return false
		}

		_, exists = ms[s[i]]
		if !exists {
			ms[s[i]] = p[i]
		}
		if ms[s[i]] != p[i] {
			return false
		}
	}

	return true
}
