package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func constructMaximumBinaryTree(nums []int) *TreeNode {
    if len(nums) == 0 {
        return nil
    }
    
    idx := 0
    for i, n := range nums {
        if n > nums[idx] {
            idx = i
        }
    }
    
    return &TreeNode {
        Val: nums[idx],
        Left: constructMaximumBinaryTree(nums[0:idx]),
        Right: constructMaximumBinaryTree(nums[idx+1:len(nums)]),
    }
}
