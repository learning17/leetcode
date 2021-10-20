package main
// https://www.nowcoder.com/practice/6ab1d9a29e88450685099d45c9e31e46

type ListNode struct{
	Val int
	Next *ListNode
}

func FindFirstCommonNode( pHead1 *ListNode ,  pHead2 *ListNode ) *ListNode {
	if pHead1 == nil || pHead2 == nil {
		return nil
	}
	p1, p2 := pHead1, pHead2
	for p1 != p2 {
		p1, p2 = p1.Next, p2.Next
		if p1 != nil && p2 == nil {
			p2 = pHead1
		}
		if p1 == nil && p2 != nil {
			p1 = pHead2
		}
	}
	return p1
}

