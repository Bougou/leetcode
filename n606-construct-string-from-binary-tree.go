package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func tree2str(t *TreeNode) string {
	if t == nil {
		return ""
	}

	var left string
	if t.Left != nil || t.Right != nil {
		left = fmt.Sprintf("(%s)", tree2str(t.Left))
	} else {
		left = ""
	}

	var right string
	if t.Right != nil {
		right = fmt.Sprintf("(%s)", tree2str(t.Right))
	} else {
		right = ""
	}

	return fmt.Sprintf("%d%s%s", t.Val, left, right)
}
