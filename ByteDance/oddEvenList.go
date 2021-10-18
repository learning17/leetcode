package main
// https://www.nowcoder.com/practice/02bf49ea45cd486daa031614f9bd6fc3

type ListNode struct{
	Val int
	Next *ListNode
}

func oddEvenList( head *ListNode ) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	odd, even := &ListNode{head.Val, nil}, &ListNode{head.Next.Val, nil}
	odd_p, even_p := odd, even
	for i, p :=3, head.Next.Next ; p != nil; p, i = p.Next, i+1 {
		node := &ListNode{p.Val, nil}
		if i %2 == 1 {
			odd_p.Next, odd_p = node, node 
		} else {
			even_p.Next, even_p = node, node
		}
	}
	odd_p.Next = even
	return odd
}
