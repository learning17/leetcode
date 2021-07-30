package main

//https://leetcode-cn.com/problems/reverse-linked-list/

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var pre, cur *ListNode
	pre, cur = nil, head
	for ; cur != nil; {
		pre, cur, cur.Next = cur, cur.Next, pre
	}
	return pre
}
