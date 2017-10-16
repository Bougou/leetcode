package main

// palindromic 具有回文结构
// palindrome 回文〔指顺读和倒读都一样的词或短语，如deed或level〕
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	var res string

	// 依次遍历每个字符，将该字符作为 palindrome 的中心，
	// 向外（两边）寻找更多的palindrome string
	for i := range []byte(s) {
		sub := ""

		// case "aba", palindrom 的中心是一个字符
		sub = findPalindrome(s, i, i)
		if len(sub) > len(res) {
			res = sub
		}

		// case "abba"，palindrome 的中心是两个字符
		sub = findPalindrome(s, i, i+1)
		if len(sub) > len(res) {
			res = sub
		}
	}

	return res
}

func findPalindrome(s string, l int, r int) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}

	return s[l+1 : r]
}
