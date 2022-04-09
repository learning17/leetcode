package main
// https://leetcode-cn.com/problems/reverse-linked-list-ii/

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	var root, leftP, rightP *ListNode
	for i, node := 1, head; ; i++ {
		if i == left - 1 {
			leftP = node
		}
		if i == left {
			root = node
		}
		if i == right {
			node.Next, rightP = nil, node.Next
			break
		}
		node = node.Next
	}
	tmpHead, tmpTail := reverse(root)
	if leftP != nil {
		leftP.Next = tmpHead
	} else {
		head = tmpHead
	}
	tmpTail.Next = rightP
	return head
}

func reverse(root *ListNode)(*ListNode, *ListNode) {
	var pre, cur *ListNode
	for pre, cur = nil,root; cur != nil ; {
		pre, cur, cur.Next = cur, cur.Next, pre
	}
	return pre, root
}

