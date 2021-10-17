package main
// https://www.nowcoder.com/practice/f23604257af94d939848729b1a5cda08

type ListNode struct{
	Val int
	Next *ListNode
}

func sortInList( head *ListNode ) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	pre, fast, slow := head, head, head
	for ;slow != nil && fast != nil && fast.Next != nil; {
		pre, slow, fast = slow, slow.Next, fast.Next.Next
	}
	pre.Next = nil
	left := sortInList(head)
	right := sortInList(slow)
	return merge(left, right)
}

func merge( pHead1 *ListNode ,  pHead2 *ListNode ) *ListNode {
    root := &ListNode{}
    node := root
    for ;pHead1 != nil && pHead2 != nil; node = node.Next {
        if pHead1.Val < pHead2.Val {
            node.Next, pHead1 = pHead1, pHead1.Next
        } else {
            node.Next, pHead2 = pHead2, pHead2.Next
        }
    }   
    if pHead1 != nil {
        node.Next = pHead1
    }   
    if pHead2 != nil {
        node.Next = pHead2                                                                                           
    }   
    return root.Next 
}

