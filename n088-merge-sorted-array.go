package main

// Put elements from end(big) to start(small)
func merge(nums1 []int, m int, nums2 []int, n int)  {
    for i := m + n - 1;  i >= 0 && m > 0 && n > 0; i-- {
        if nums1[m-1] > nums2[n-1] {
            nums1[i] = nums1[m-1]
            m--
        } else {
            nums1[i] = nums2[n-1]
            n--
        }
    }
    
    //If m firstly reaches to 0, put all elements left in nums2 to num1
    if m == 0 {
        for ; n > 0; n-- {
            nums1[n-1] = nums2[n-1]
        }
    }
}
