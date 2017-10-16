package main

// Not the question is to find the most left of the most depth leaf level.
// The node does not have to be a left child.
func findBottomLeftValue(root *TreeNode) int {
	queue := []*TreeNode{root}

	var node *TreeNode

	for len(queue) != 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			node = queue[0]

			if l > 1 {
				queue = queue[1:len(queue)]
			} else if l == 1 {
				queue = queue[0:0]
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
		}
	}

	return node.Val
}
