// Given a binary tree, determine if it is height-balanced.

// For this problem, a height-balanced binary tree is defined as a binary tree in which the depth of the two subtrees of every node never differ by more than 1.

package main

// the depth of the tree
func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	ld := depth(root.Left)
	rd := depth(root.Right)

	if ld > rd {
		return ld + 1
	}

	return rd + 1
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	ld := depth(root.Left)
	rd := depth(root.Right)

	d := ld - rd
	if d < 0 {
		d = -d
	}

	return d <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

// Another Solution: O(n)

// if the tree is not balanced, return -1, otherwise return the depth
// when the depth of node's left or node's right is -1, then the depth of the node is -1
func depthOrNeg(root *TreeNode) int {
	if root == nil {
		return 0
	}

	ld := depthOrNeg(root.Left)
	rd := depthOrNeg(root.Right)

	if ld == -1 || rd == -1 {
		return -1
	}

	d := ld - rd
	if d > 1 || d < -1 {
		return -1
	}

	if ld > rd {
		return ld + 1
	}

	return rd + 1
}

func isBalanced2(root *TreeNode) bool {
	return depthOrNeg(root) != -1
}
