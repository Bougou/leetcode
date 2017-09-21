package main

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	var tree TreeNode

	if t1 == nil && t2 == nil {
		return nil
	}

	if t1 != nil && t2 == nil {
		return t1
	} else if t1 == nil && t2 != nil {
		return t2
	} else {
		tree.Val = t1.Val + t2.Val
		tree.Left = mergeTrees(t1.Left, t2.Left)
		tree.Right = mergeTrees(t1.Right, t2.Right)
	}

	return &tree
}
