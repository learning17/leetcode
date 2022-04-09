package main

// https://leetcode-cn.com/problems/reverse-linked-list-ii/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	var pre, subRoot *ListNode
	node := head
	for i := 1; i < left; i++ {
		pre, node = node, node.Next
	}
	subRoot = node

	var subTail, next *ListNode
	node = head
	for i := 0; i < right; i++ {
		subTail, node = node, node.Next
	}
	next = node
	if pre != nil {
		pre.Next = nil
	}
	subTail.Next = nil
	subRoot, subTail = reverse(subRoot)
	if pre != nil {
		pre.Next = subRoot
	} else {
		head = subRoot
	}
	subTail.Next = next
	return head
}

func reverse(head *ListNode) (*ListNode, *ListNode) {
	var pre, cur *ListNode
	for pre, cur = nil, head; cur != nil; {
		pre, cur, cur.Next = cur, cur.Next, pre
	}
	return pre, head
}
