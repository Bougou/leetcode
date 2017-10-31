// Given a binary tree, return the level order traversal of its nodes' values. (ie, from left to right, level by level).

// For example:
// Given binary tree [3,9,20,null,null,15,7],
//     3
//    / \
//   9  20
//     /  \
//    15   7
// return its level order traversal as:
// [
//   [3],
//   [9,20],
//   [15,7]
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
func levelOrder(root *TreeNode) [][]int {
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
		for _, node := range allNodes[len(allNodes)-1] {
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
			allNodes = append(allNodes, nodes)
			allVals = append(allVals, vals)
		} else {
			break
		}
	}

	return allVals
}
