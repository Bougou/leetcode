// Given a singly linked list where elements are sorted in ascending order, convert it to a height balanced BST.

package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	// len of the list
	l := 0
	for p := head; p != nil; {
		l++
		p = p.Next
	}

	// find the middle node in the list, and the pre and post node of middle
	var pre *ListNode
	var post *ListNode
	mid := head
	post = head.Next
	for i := 0; i < l/2; i++ {
		pre = mid
		mid = mid.Next
		post = mid.Next
	}
	// break the list
	if pre != nil {
		pre.Next = nil
	}

	root := &TreeNode{
		Val: mid.Val,
	}

	if pre != nil {
		root.Left = sortedListToBST(head)
	}

	if post != nil {
		root.Right = sortedListToBST(post)
	}

	return root
}
