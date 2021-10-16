package main
// https://www.nowcoder.com/practice/b49c3dc907814e9bbfa8437c251b028e

type ListNode struct{
	Val int
	Next *ListNode
}

func reverseKGroup( head *ListNode ,  k int ) *ListNode {
	var point,tmpPoint, rHead, rTail *ListNode
	for {
		if head == nil {
			break
		}
		rHead, rTail, head = reverse(head, k)
		if point == nil {
			point, tmpPoint = rHead, rTail
		} else {
			tmpPoint.Next, tmpPoint = rHead, rTail
		}
	}
	return point
}

func reverse(head *ListNode, k int)(*ListNode, *ListNode, *ListNode) {
	var pre, cur *ListNode
	var i int
	for pre, cur = nil, head; i < k && cur != nil; {
		pre, cur, cur.Next = cur, cur.Next, pre
		i++
	}
	if i < k {
		head = pre
		for pre, cur = nil, head; cur != nil; {
			pre, cur, cur.Next = cur, cur.Next, pre
		}
	}
	return pre, head, cur
}


