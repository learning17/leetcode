package main
// https://leetcode-cn.com/problems/7WHec2/

type ListNode struct {
	Val int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	pre, fast, slow := head,head, head
	for;slow != nil && fast != nil && fast.Next != nil; {
		pre, slow, fast = slow, slow.Next, fast.Next.Next
	}
	pre.Next = nil
	return merge(sortList(head), sortList(slow))
}

func merge(left, right *ListNode) *ListNode {
	root := &ListNode{}
	p := root
	for ;left != nil && right != nil; {
		if left.Val < right.Val {
			p.Next, left = left, left.Next
		} else {
			p.Next, right = right, right.Next
		}
		p = p.Next
	}
	if left != nil {
		p.Next = left 
	}
	if right != nil {
		p.Next = right
	}
	return root.Next
}


