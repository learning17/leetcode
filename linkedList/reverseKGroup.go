package main
// https://leetcode-cn.com/problems/reverse-nodes-in-k-group/
// K 个一组翻转链表

type ListNode struct {
	Val int
	Next *ListNode
}
func reverseKGroup(head *ListNode, k int) *ListNode {
	var root, tail *ListNode
	var tmpHead, tmpTail, tmpNext *ListNode
	for root, tmpNext = nil, head; tmpNext != nil; {
		tmpHead, tmpTail, tmpNext = reverseK(tmpNext, k)
		if root == nil {
			root, tail = tmpHead, tmpTail
		} else {
			tail.Next, tail = tmpHead, tmpTail
		}
	}
	return root
}

func reverseK(root *ListNode, k int) (*ListNode, *ListNode, *ListNode) {
	var pre, cur *ListNode
	var i int
	for pre, cur, i = nil, root, 0; i < k && cur != nil; i++ {
		pre, cur, cur.Next = cur, cur.Next, pre
	}
	if i < k {
		return reverseK(pre, i)
	}
	return pre, root, cur
}
