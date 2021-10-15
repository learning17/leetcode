package main
// https://www.nowcoder.com/practice/c56f6c70fb3f4849bc56e33ff2a50b6b

type ListNode struct{
	Val int
	Next *ListNode
}

func addInList( head1 *ListNode ,  head2 *ListNode ) *ListNode {
	head1, head2 = reverse(head1), reverse(head2)
	node1, node2 := head1, head2
	pre1 := head1
	bit := 0
	for ;node1 != nil && node2 != nil; pre1, node1, node2 = node1,node1.Next, node2.Next {
		value := node1.Val + node2.Val + bit
		node1.Val, bit = value%10, value/10
	}
	if node1 == nil {
		pre1.Next, node1 = node2, node2
	}
	for ;node1 != nil; node1 = node1.Next {
		if bit == 0 {
			break
		}
		value := node1.Val + bit
		node1.Val, bit = value%10, value/10
	}
	head1 = reverse(head1)
	if bit != 0 {
		head1 = &ListNode{
			Val: bit,
			Next: head1,
		}
	}
	return head1
}

func reverse(head *ListNode) *ListNode {
	var pre, cur *ListNode
	for pre, cur = nil, head; cur != nil; {
		pre, cur, cur.Next = cur, cur.Next, pre
	}
	return pre
}
