package main

func lengthOfLongestSubstring(s string) int {
	ret := 0
	curSub := []byte{}

	for _, c := range []byte(s) {
		flag := 0
		index := 0
		for i, v := range curSub {
			if c == v {
				flag = 1
				index = i
				break
			}
		}

		if flag == 1 {
			if len(curSub) > ret {
				ret = len(curSub)
			}
			curSub = curSub[index+1 : len(curSub)]
		}
		curSub = append(curSub, c)
	}

	// for the last substring sequence
	if len(curSub) > ret {
		ret = len(curSub)
	}

	return ret
}
