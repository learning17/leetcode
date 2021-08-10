package main

// https://leetcode-cn.com/problems/reorder-list/

type ListNode struct {
	Val int
	Next *ListNode
}

func reorderList(head *ListNode)  {

}

func reverList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre, cur *ListNode
	pre, cur = nil, head
	for ;cur != nil; {
		pre, cur, cur.Next = cur, cur.Next, pre
	}
}
