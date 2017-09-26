package main

func canConstruct(ransomNote string, magazine string) bool {
	a := [26]int{}

	for _, v := range []rune(magazine) {
		index := int(v - 'a')
		a[index]++
	}

	for _, v := range []rune(ransomNote) {
		index := int(v - 'a')
		a[index]--
		if a[index] < 0 {
			return false
		}
	}

	return true
}
