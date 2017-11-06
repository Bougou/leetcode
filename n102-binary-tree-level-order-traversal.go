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
	
	resNode := [][]*TreeNode{
			[]*TreeNode{root},
	}
	
	resVal := [][]int{
			[]int{root.Val},
	}
	
	for i:= 0; ; i++{ 
			nextLevelNode := []*TreeNode{}
			nextLevelVal := []int{}
			nextLevelCount := 0
			for _, node := range resNode[i] {
					if node.Left != nil {
							nextLevelNode = append(nextLevelNode, node.Left)
							nextLevelVal = append(nextLevelVal, node.Left.Val)
							nextLevelCount++
					}
					
					if node.Right != nil {
							nextLevelNode = append(nextLevelNode, node.Right)
							nextLevelVal = append(nextLevelVal, node.Right.Val)
							nextLevelCount++
					}
			}
			
			if nextLevelCount > 0 {
					resNode = append(resNode, nextLevelNode)
					resVal = append(resVal, nextLevelVal)
			} else {
					break
			}
	}
	
	return resVal
}