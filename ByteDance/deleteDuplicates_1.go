package main
// https://www.nowcoder.com/practice/c087914fae584da886a0091e877f2c79
type ListNode struct{
	Val int
	Next *ListNode
}

func deleteDuplicates( head *ListNode ) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre, cur *ListNode
	for pre, cur = head, head.Next; cur != nil; {
		if pre.Val != cur.Val {
			pre, cur = cur, cur.Next
		} else {
			for ;cur != nil && cur.Val == pre.Val; cur = cur.Next {
			}
			pre.Next = cur
		}
	}
	return head
}

