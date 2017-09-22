package main

import (
	"fmt"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/**
 * 构造一个map变量
 * 在遍历每一个节点时，将Val值以键存到一个map变量中
 * 在遍历每一个节点时，判断 k - Val 的差值是否存在于map的键中，
 * 如果找到了相应的键，说明树中存在两个节点的值相加等于k（当前遍历的值 和map中存在的变量）
 */

func findTarget(root *TreeNode, k int) bool {
	dataMap := make(map[int]int)

	return dfs(root, k, dataMap)
}

func dfs(root *TreeNode, k int, data map[int]int) bool {
	if root == nil {
		return false
	}

	_, exists := data[k-root.Val]

	if exists {
		return true
	}

	data[root.Val] = 1

	fmt.Println(data)
	return dfs(root.Left, k, data) || dfs(root.Right, k, data)
}
