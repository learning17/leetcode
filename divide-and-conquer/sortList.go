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
	pre, slow, fast := head, head, head
	for slow != nil && fast != nil && fast.Next != nil {
		pre, slow, fast = slow, slow.Next, fast.Next.Next
	}
	pre.Next = nil
	return merge(sortList(head), sortList(slow))
}

func merge(root1, root2 *ListNode) *ListNode {
	root := &ListNode{}
	node := root
	for ;root1 != nil && root2 != nil; node = node.Next {
		if root1.Val < root2.Val {
			node.Next, root1 = root1, root1.Next
		} else {
			node.Next, root2 = root2, root2.Next
		}
	}
	if root1 != nil {
		node.Next = root1
	}
	if root2 != nil {
		node.Next = root2
	}
	return root.Next
}

