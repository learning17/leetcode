package main
// https://www.nowcoder.com/practice/65cfde9e5b9b4cf2b6bafa5f3ef33fa6
type ListNode struct{
	Val int
	Next *ListNode
}

func mergeKLists( lists []*ListNode ) *ListNode {
	size := len(lists)
	if size == 0 {
		return nil
	}
	if size == 1 {
		return lists[0]
	}
	mid := size/2
	left := mergeKLists(lists[:mid])
	right := mergeKLists(lists[mid:size])
	root := &ListNode{}
	node := root
	for left != nil && right != nil {
		if left.Val < right.Val {
			node.Next, left = left, left.Next
		} else {
			node.Next, right = right, right.Next
		}
		node = node.Next
	}
	if left != nil {
		node.Next = left
	}
	if right != nil {
		node.Next = right
	}
	return root.Next
}

