package main
// https://leetcode-cn.com/problems/linked-list-cycle-lcci/

type ListNode struct {
	Val int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	for {
		if slow == nil || fast == nil || fast.Next == nil {
			return nil
		}
		fast, slow = fast.Next.Next, slow.Next
		if fast == slow && slow != nil {
			break
		}
	}

	fast = head
	for {
		if fast == slow {
			return slow
		}
		fast, slow = fast.Next, slow.Next
	}
}
