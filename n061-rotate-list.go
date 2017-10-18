package main

import "fmt"

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var cur *ListNode
	len := 1
	for cur = head; cur.Next != nil; cur = cur.Next {
		len++
	}

	fmt.Println(len)
	m := k % len
	if m == 0 {
		return head
	}

	// circle the list
	cur.Next = head
	var newHead *ListNode
	for i := 1; i < len-m; i++ {
		head = head.Next
	}
	newHead = head.Next
	head.Next = nil

	return newHead
}

func iterateList(head *ListNode) []int {
	res := []int{head.Val}
	for head.Next != nil {
		res = append(res, head.Next.Val)
		head = head.Next
	}

	return res
}
