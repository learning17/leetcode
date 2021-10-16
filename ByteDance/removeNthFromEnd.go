package main
// https://www.nowcoder.com/practice/f95dcdafbde44b22a6d741baf71653f6

type ListNode struct{
	Val int
	Next *ListNode
}

func removeNthFromEnd( head *ListNode ,  n int ) *ListNode {
	size := 0
	for p := head; p != nil; p = p.Next {
		size++
	}
	if size == n {
		return head.Next
	}
	k := size - n
	for pre, cur := head, head; cur != nil && k >=0; k-- {
		if k == 0 {
			pre.Next = cur.Next
		} else {
			pre, cur = cur, cur.Next
		}
	}
	return head
}
