package main

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxSubLeft := maxDepth(root.Left)
	maxSubRight := maxDepth(root.Right)

	if maxSubLeft > maxSubRight {
		return 1 + maxSubLeft
	}

	return 1 + maxSubRight
}
