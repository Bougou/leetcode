// Given a binary tree, find its minimum depth.

// The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.

package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	ld := minDepth(root.Left)
	rd := minDepth(root.Right)

	if ld == 0 {
		return rd + 1
	}

	if rd == 0 {
		return ld + 1
	}

	if ld < rd {
		return ld + 1
	}

	return rd + 1
}
