package main
// https://leetcode-cn.com/problems/aMhZSa/
// 回文链表

type ListNode struct {
	Val int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	pre, slow, fast := head, head, head
	for slow != nil && fast != nil && fast.Next != nil {
		pre, slow, fast = slow, slow.Next, fast.Next.Next
	}
	pre.Next = nil
	slow = reverse(slow)
	for ;head != nil && slow != nil; head, slow = head.Next, slow.Next {
		if head.Val != slow.Val {
			return false
		}
	} 
	return true
}

func reverse(root *ListNode) *ListNode {
	var pre, cur *ListNode
	for pre, cur = nil, root; cur != nil; {
		pre, cur, cur.Next = cur, cur.Next, pre
	}
	return pre
}
