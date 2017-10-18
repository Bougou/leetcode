package main

// Interleaving does not mean that the segments must be the same length, only that by advancing in one string at a time, we can recreate s3.

// https://discuss.leetcode.com/topic/897/definition-of-interleave-on-test-examples
// https://github.com/haoel/leetcode/blob/2f3a77de45a6addf7aae3fb6988e8422a7a3eb6c/algorithms/cpp/interleavingString/interleavingString.cpp

// Dynamic Programming
func isInterleave(s1 string, s2 string, s3 string) bool {
	l1 := len(s1)
	l2 := len(s2)
	l3 := len(s3)

	if l1+l2 != l3 {
		return false
	}

	if l1 == 0 {
		if s2 == s3 {
			return true
		}
		return false
	}

	if l2 == 0 {
		if s1 == s3 {
			return true
		}
		return false
	}

	var m = make([][]bool, l1+1)
	for i := range m {
		m[i] = make([]bool, l2+1)
	}

	m[0][0] = true

	// the top line
	for i := 1; i <= l2; i++ {
		if s2[i-1] == s3[i-1] {
			m[0][i] = true
		} else {
			break
			// leave m[0][i] false default
		}
	}

	// the left line
	for i := 1; i <= l1; i++ {
		if s1[i-1] == s3[i-1] {
			m[i][0] = true
		} else {
			break
			// leave m[i][0] false default
		}
	}

	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if s3[i+j-1] == s1[i-1] {
				m[i][j] = m[i-1][j] || m[i][j]
			}

			if s3[i+j-1] == s2[j-1] {
				m[i][j] = m[i][j-1] || m[i][j]
			}
		}
	}

	return m[l1][l2]
}
