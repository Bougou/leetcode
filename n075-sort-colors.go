package main

import "fmt"

func sortColors(nums []int) {
	c2 := 0
	i := 0
	for i < len(nums)-c2 {
		if nums[i] == 0 {
			nums = append([]int{0}, append(nums[:i], nums[i+1:]...)...)
			i++
		} else if nums[i] == 2 {
			nums = append(append(nums[:i], nums[i+1:]...), 2)
			c2++
		} else {
			i++
		}
		fmt.Println(nums)
	}
	fmt.Println(nums)
}
