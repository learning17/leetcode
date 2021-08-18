package main

// https://leetcode-cn.com/problems/linked-list-cycle/

type ListNode struct {
	Val int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow, fast = slow.Next, fast.Next.Next
		if slow == fast {
			return true
		}
	}
}
