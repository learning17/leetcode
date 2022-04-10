package main
// https://leetcode-cn.com/problems/merge-k-sorted-lists/

type ListNode struct {
	Val int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	size := len(lists)
	if size == 0 {
		return nil
	}
	return merge(0, size-1, lists)
}

func merge(left, right int, lists []*ListNode) *ListNode {
	if left >= right {
		return lists[left]
	}
	mid := left + (right - left)/2
	leftP := merge(left, mid, lists)
	rightP := merge(mid+1, right, lists)
	return merge2Lists(leftP, rightP)
}

func merge2Lists(root1 *ListNode, root2 *ListNode) *ListNode {
	root := &ListNode{}
	node := root
	for;root1 != nil && root2 != nil; {
		if root1.Val < root2.Val {
			node.Next, root1 = root1, root1.Next
		} else {
			node.Next, root2 = root2, root2.Next
		}
		node = node.Next
	}
	if root1 != nil {
		node.Next = root1
	}
	if root2 != nil {
		node.Next = root2
	}
	return root.Next
}
