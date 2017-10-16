package main

import "math"

func findDisappearedNumbers(nums []int) []int {
	// 遍历数组，以每个元素的绝对值减去1作为索引, 将数组的该索引位置元素置为负数
	// 之后再次遍历该数组，过滤出元素值依然为正数的数组索引，那么该索引位置+1即为源数组中没有出现过的数字
	for i := 0; i < len(nums); i++ {
		ele := nums[i]
		index := int(math.Abs(float64(ele)) - 1)
		nums[index] = -int(math.Abs(float64(nums[index])))
	}

	ret := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			ret = append(ret, i+1)
		}
	}

	return ret
}
