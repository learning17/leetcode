package main
// https://www.nowcoder.com/practice/71cef9f8b5564579bf7ed93fbe0b2024

type ListNode struct{
	Val int
	Next *ListNode
}

func deleteDuplicates( head *ListNode ) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre, cur *ListNode
	pre, cur = nil, head
	for cur != nil && cur.Next != nil {
		if cur.Val != cur.Next.Val {
			pre, cur = cur, cur.Next
		} else {
			point := cur
			for ;cur != nil && cur.Val == point.Val; cur = cur.Next{
			}
			if pre != nil {
				pre.Next = cur
			} else {
				head = cur
			}
		}
	}
	return head
}
