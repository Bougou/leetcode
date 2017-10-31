// Given an array where elements are sorted in ascending order, convert it to a height balanced BST.

package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	mid := len(nums) / 2
	root := &TreeNode{
		Val: nums[mid],
	}

	if mid > 0 {
		root.Left = sortedArrayToBST(nums[:mid])
	}

	if mid+1 < len(nums) {
		root.Right = sortedArrayToBST(nums[mid+1:])
	}

	return root
}
