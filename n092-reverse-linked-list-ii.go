// Reverse a linked list from position m to n. Do it in-place and in one-pass.

// For example:
// Given 1->2->3->4->5->NULL, m = 2 and n = 4,

// return 1->4->3->2->5->NULL.

// Note:
// Given m, n satisfy the following condition:
// 1 ≤ m ≤ n ≤ length of list.

package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{}
	dummy.Next = head

	pre := dummy
	for i := 0; i < m-1; i++ {
		pre = pre.Next
	}
	// now `pre` points to the node before reverse

	start := pre.Next
	then := start.Next

	for i := 0; i < n-m; i++ {
		start.Next = then.Next
		then.Next = pre.Next
		pre.Next = then

		then = start.Next
	}

	return dummy.Next
}

// 1->2->3->4->5->6
// m=3, n=6
// dummy->1->2->3->4->5->6
// Find out pre: dummy->1->2(pre)->3(start)->4(then)->5->6
// 1 loop: dummy->1->2(pre)->4->3(start)->5(then)->6
// 2 loop: dummy->1->2(pre)->5->4->3(start)->6(then)
// 3 loop: dummy->1->2(pre)->6->5->4->3(start)->null(then)
// return 1->2->6->5->4->3
