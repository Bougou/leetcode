// Given a binary tree, return the bottom-up level order traversal of its nodes' values. (ie, from left to right, level by level from leaf to root).

// For example:
// Given binary tree [3,9,20,null,null,15,7],
//     3
//    / \
//   9  20
//     /  \
//    15   7
// return its bottom-up level order traversal as:
// [
//   [15,7],
//   [9,20],
//   [3]
// ]

package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	allNodes := [][]*TreeNode{
		[]*TreeNode{root},
	}

	allVals := [][]int{
		[]int{root.Val},
	}

	for {
		nodes := []*TreeNode{}
		vals := []int{}
		for _, node := range allNodes[0] {
			if node.Left != nil {
				nodes = append(nodes, node.Left)
				vals = append(vals, node.Left.Val)

			}
			if node.Right != nil {
				nodes = append(nodes, node.Right)
				vals = append(vals, node.Right.Val)
			}
		}

		if len(nodes) > 0 {
			allNodes = append([][]*TreeNode{nodes}, allNodes...)
			allVals = append([][]int{vals}, allVals...)
		} else {
			break
		}
	}

	return allVals
}
