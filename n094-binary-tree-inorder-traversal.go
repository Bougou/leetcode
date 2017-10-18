package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Recursive Solution
func inorderTraversal2(root *TreeNode) []int {
    res := []int{}
    if root == nil {
        return res
    }
    
    left := inorderTraversal(root.Left)
    right := inorderTraversal(root.Right)
    
    res = append(res, left...)
    res = append(res, root.Val)
    res = append(res, right...)
    return res
}

// Iterative Solution
// Use a stack-like data structure
func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	stack := []*TreeNode{}

	cur := root
	for cur != nil || len(stack) != 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		cur = stack[len(stack)-1]
		stack = stack[0:len(stack)-1]
		res = append(res, cur.Val)
		cur = cur.Right
	}

	return res
}
