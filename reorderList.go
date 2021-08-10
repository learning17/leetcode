package main

// https://leetcode-cn.com/problems/reorder-list/

type ListNode struct {
	Val int
	Next *ListNode
}

func reorderList(head *ListNode)  {
	slow, fast := head, head
	for; fast != nil && fast.Next != nil; {
		slow, fast = slow.Next, fast.Next.Next
	}
	if fast != nil {
		slow = slow.Next
	}
	first, second := head, reverList(slow)
	for ;first != nil && second != nil; {
		first.Next, second.Next, first, second = second, first.Next, first.Next, second.Next
	}
	if first != nil {
		first.Next = nil
	}
}

func reverList(head *ListNode) *ListNode {
	var pre, cur *ListNode
	pre, cur = nil, head
	for ;cur != nil; {
		pre, cur, cur.Next = cur, cur.Next, pre
	}
	return pre
}
