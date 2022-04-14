package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var pre, slow, fast *ListNode
	pre, slow, fast = nil, head, head
	for fast != nil && fast.Next != nil {
		pre, slow, fast = slow, slow.Next, fast.Next.Next
	}
	pre.Next = nil
	left, right := sortList(head), sortList(slow)
	return merge(left, right)
}

func merge(list1 *ListNode, list2 *ListNode) *ListNode {
	root := &ListNode{}
	node := root
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			node.Next, list1 = list1, list1.Next
		} else {
			node.Next, list2 = list2, list2.Next
		}
		node = node.Next
	}
	if list2 != nil {
		list1 = list2
	}
	node.Next = list1
	return root.Next
}
