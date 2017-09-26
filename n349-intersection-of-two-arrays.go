package main

func intersection(nums1 []int, nums2 []int) []int {
	data := map[int]int{}
	ret := []int{}

	for _, num1 := range nums1 {
		for _, num2 := range nums2 {
			if num2 == num1 {
				if _, exists := data[num1]; !exists {
					ret = append(ret, num1)
					data[num1] = 1
				}
			}
		}
	}

	return ret
}
