package main
// https://www.nowcoder.com/practice/b58434e200a648c589ca2063f1faf58c

type ListNode struct{
	Val int
	Next *ListNode
}

func reverseBetween( head *ListNode ,  m int ,  n int ) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var p1,p2, pre, cur *ListNode
	p1,p2, pre, cur = nil, nil, nil, head
	for i := 1; cur != nil;i++ {
		if i < m {
			p1, cur = cur, cur.Next
		} else if i >= m && i <= n {
			if p2 == nil {
				p2 = cur
			}
			pre, cur, cur.Next = cur, cur.Next, pre
		}
	}
	if p1 != nil {
		p1.Next = pre
	} else {
		head = pre
	}
	if p2 != nil {
		p2.Next = cur
	}
	return head
}

