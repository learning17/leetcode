package main
// https://leetcode-cn.com/problems/lMSNwu/
// 链表中的两数相加

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1, l2 = reverse(l1), reverse(l2)
	root := &ListNode{}
	node := root
	var bit int
	for ;l1 != nil && l2 != nil; l1, l2, node = l1.Next, l2.Next, node.Next {
		sum := l1.Val + l2.Val + bit
		bit = sum / 10
		node.Next = &ListNode{sum % 10, nil}
	}
	if l1 == nil {
		l1 = l2
	}
	for ;l1 != nil; l1, node = l1.Next, node.Next {
		sum := l1.Val + bit
		bit = sum / 10
		node.Next = &ListNode{sum % 10, nil}
	}
	if bit != 0 {
		node.Next = &ListNode{bit, nil}
	}
	return reverse(root.Next)
}

func reverse(root *ListNode) *ListNode {
	var pre, cur *ListNode
	for pre, cur = nil, root; cur != nil; {
		pre, cur, cur.Next = cur, cur.Next, pre 
	}
	return pre
}

