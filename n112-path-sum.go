package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	stack := []*TreeNode{root}
	preSum := []int{0}

	s := 0
	for {
		l := len(stack)
		if s >= l {
			break
		}

		for i := s; i < l; i++ {
			nodeSum := stack[i].Val + preSum[i]
			if nodeSum == sum && stack[i].Left == nil && stack[i].Right == nil {
				return true
			}

			if stack[i].Left != nil {
				stack = append(stack, stack[i].Left)
				preSum = append(preSum, nodeSum)
			}

			if stack[i].Right != nil {
				stack = append(stack, stack[i].Right)
				preSum = append(preSum, nodeSum)
			}
		}
		s = l
	}

	return false
}
