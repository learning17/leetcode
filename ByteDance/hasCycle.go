package main
// https://www.nowcoder.com/practice/650474f313294468a4ded3ce0f7898b9

type ListNode struct{
	Val int
	Next *ListNode
}

func hasCycle( head *ListNode ) bool {
	slow, fast := head, head
	for ;slow != nil && fast != nil && fast.Next != nil; {
		slow, fast = slow.Next, fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
