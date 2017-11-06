// Given an index k, return the kth row of the Pascal's triangle.

// For example, given k = 3,
// Return [1,3,3,1].

// Note:
// Could you optimize your algorithm to use only O(k) extra space?

package main

func getRow(rowIndex int) []int {
	res := []int{1}
	for i := 1; i <= rowIndex; i++ {
		res = append(res, 1)
		for j := 1; j <= i/2; j++ {
			tmp := res[j] + res[len(res)-j-1]
			res[j] = tmp
			res[len(res)-j-1] = tmp
		}
	}

	return res
}
