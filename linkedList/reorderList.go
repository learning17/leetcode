package main
// https://leetcode-cn.com/problems/LGjMqU/
// 重排链表

type ListNode struct {
	Val int
	Next *ListNode
}

func reorderList(head *ListNode)  {
	if head == nil || head.Next == nil {
        return
    }
	pre, slow, fast := head, head, head
	for slow != nil && fast != nil && fast.Next != nil {
		pre, slow, fast = slow, slow.Next, fast.Next.Next
	}
	pre.Next = nil
	slow = reverse(slow)
	root := &ListNode{}
	node := root
	for ;head != nil && slow != nil; {
		node.Next, node, head = head, head, head.Next
		node.Next, node, slow = slow, slow, slow.Next
	}
	if head != nil {
		node.Next = head
	}
	head = root.Next
}

func reverse(root *ListNode) *ListNode {
	var pre, cur *ListNode
	for pre, cur = nil, root; cur != nil; {
		pre, cur, cur.Next = cur, cur.Next, pre
	}
	return pre
}
