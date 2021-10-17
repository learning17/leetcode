package main
//https://www.nowcoder.com/practice/3fed228444e740c8be66232ce8b87c2f

type ListNode struct{
	Val int
	Next *ListNode
}

func isPail( head *ListNode ) bool {
	if head == nil || head.Next == nil {
		return true
	}
	pre, fast, slow := head, head, head
	for ;slow != nil && fast != nil && fast.Next != nil; {
		pre, slow, fast = slow, slow.Next, fast.Next.Next
	}
	pre.Next = nil
	cur := slow
	for cur != nil {
		pre, cur, cur.Next = cur, cur.Next, pre
	}
	for ;pre != nil && head != nil; pre, head = pre.Next, head.Next {
		if pre.Val != head.Val {
			return false
		}
	}
	return true
}

