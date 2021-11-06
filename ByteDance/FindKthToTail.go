package main
// https://www.nowcoder.com/practice/886370fe658f41b498d40fb34ae76ff9

type ListNode struct{
	Val int
	Next *ListNode
}

func FindKthToTail( pHead *ListNode ,  k int ) *ListNode {
	fast, slow := pHead, pHead
	step := 0
	for fast != nil {
		fast = fast.Next
		if step >= k {
			slow = slow.Next
		}
		step++
	}
	if step < k {
		slow = nil
	}
	return slow
}

