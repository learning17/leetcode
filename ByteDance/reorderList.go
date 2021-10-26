package main
// https://www.nowcoder.com/practice/3d281dc0b3704347846a110bf561ef6b

type ListNode struct{
	Val int
	Next *ListNode
}

func reorderList( head *ListNode )  {
	if head == nil || head.Next == nil {
		return
	}
	pre, fast, slow := head, head, head
	for slow != nil && fast != nil && fast.Next != nil {
		pre, slow, fast = slow, slow.Next, fast.Next.Next
	}
	pre.Next = nil
	head1, head2 := head, reserve(slow)
	root := &ListNode{}
	node := root
	for ;head1 != nil && head2 != nil; {
		node.Next, head1 = head1, head1.Next
		node = node.Next
		node.Next, head2 = head2, head2.Next
		node = node.Next
	}
	head = root.Next
}

func reserve(node *ListNode) *ListNode {
	var pre, cur *ListNode
	for pre, cur = nil, node; cur != nil; {
		pre, cur, cur.Next = cur, cur.Next, pre
	}
	return pre
}

