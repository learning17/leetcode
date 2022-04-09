package main

// 反转链表
//https://leetcode-cn.com/problems/reverse-linked-list/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var cur, pre *ListNode
	for pre, cur = nil, head; cur != nil; {
		pre, cur, cur.Next = cur, cur.Next, pre
	}
	return pre
}
