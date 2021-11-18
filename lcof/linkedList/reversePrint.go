package main
// https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/
// 从尾到头打印链表

type ListNode struct {
	Val int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	head = reverse(head)
	var res []int
	for ;head != nil; head = head.Next {
		res = append(res, head.Val)
	}
	return res
}

func reverse(root *ListNode) *ListNode {
	var pre, cur *ListNode
	for pre, cur = nil, root; cur != nil; {
		pre, cur, cur.Next = cur, cur.Next, pre
	}
	return pre
}

