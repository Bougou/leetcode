package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	var ret *ListNode
	var pre *ListNode

	for l1 != nil || l2 != nil || carry != 0 {
		a := 0
		b := 0
		if l1 != nil {
			a = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			b = l2.Val
			l2 = l2.Next
		}

		sum := (a + b + carry) % 10
		carry = (a + b + carry) / 10

		cur := &ListNode{
			Val:  sum,
			Next: nil,
		}

		if pre == nil {
			// the head node
			ret = cur
		} else {
			pre.Next = cur
		}

		pre = cur
	}

	return ret
}
