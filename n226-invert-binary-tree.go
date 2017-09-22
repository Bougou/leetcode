package main

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	left := invertTree(root.Right)
	right := invertTree(root.Left)

	root.Left = left
	root.Right = right

	return root
}

func invertTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	tree := TreeNode{
		Val:   root.Val,
		Left:  invertTree(root.Right),
		Right: invertTree(root.Left),
	}

	return &tree
}
