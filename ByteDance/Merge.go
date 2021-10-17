package main
// https://www.nowcoder.com/practice/d8b6b4358f774294a89de2a6ac4d9337

type ListNode struct{
	Val int
	Next *ListNode
}

func Merge( pHead1 *ListNode ,  pHead2 *ListNode ) *ListNode {
	root := &ListNode{}
	node := root
	for ;pHead1 != nil && pHead2 != nil; node = node.Next {
		if pHead1.Val < pHead2.Val {
			node.Next, pHead1 = pHead1, pHead1.Next
		} else {
			node.Next, pHead2 = pHead2, pHead2.Next
		}
	}
	if pHead1 != nil {
		node.Next = pHead1
	}
	if pHead2 != nil {
		node.Next = pHead2
	}
	return root.Next 
}
