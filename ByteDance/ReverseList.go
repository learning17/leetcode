package main

type ListNode struct{
	Val int
	Next *ListNode
}

func ReverseList( pHead *ListNode ) *ListNode {
	var pre, cur *ListNode
	for pre, cur = nil, pHead; cur != nil; {
		pre, cur, cur.Next = cur, cur.Next, pre
	}
	return pre
}
