package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	}

	// found out the shortest length of the element in the array
	minLen := len(strs[0])
	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < minLen {
			minLen = len(strs[i])
		}
	}

	pattern := ""
	for i := 0; i < minLen; i++ {
		p := strs[0][i] // byte
		for j := 1; j < len(strs); j++ {
			if strs[j][i] != p {
				return pattern
			}
		}
		pattern = pattern + string(p)
	}

	return pattern
}
