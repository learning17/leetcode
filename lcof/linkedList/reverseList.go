package main
// https://leetcode-cn.com/problems/fan-zhuan-lian-biao-lcof/
// 反转链表

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var pre, cur *ListNode
	for pre, cur = nil, head; cur != nil; {
		pre, cur, cur.Next = cur, cur.Next, pre
	}
	return pre
}
