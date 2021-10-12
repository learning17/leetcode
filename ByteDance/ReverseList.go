package main
//https://www.nowcoder.com/practice/75e878df47f24fdc9dc3e400ec6058ca

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
