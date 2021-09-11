package main
// https://leetcode-cn.com/problems/vvXgSW/

type ListNode struct {
	Val int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	size := len(lists)
	if size == 0 {
		return nil
	}
	if size == 1 {
		return lists[0]
	}
	mid := size / 2
	return merge(mergeKLists(lists[:mid]), mergeKLists(lists[mid:]))
}

func merge(first, second *ListNode) *ListNode {
	root := &ListNode{}
	p := root
	for ;first != nil && second != nil; {
		if first.Val < second.Val {
			p.Next, first = first, first.Next
		} else {
			p.Next, second = second, second.Next
		}
		p = p.Next
	}
	if first != nil {
		p.Next = first
	}
	if second != nil {
		p.Next = second
	}
	return root.Next
}
