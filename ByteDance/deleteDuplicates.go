package main
// https://www.nowcoder.com/practice/71cef9f8b5564579bf7ed93fbe0b2024

type ListNode struct{
	Val int
	Next *ListNode
}

func deleteDuplicates( head *ListNode ) *ListNode {
    if head == nil {
        return nil
    }
    var pre, cur *ListNode
    for pre, cur = nil, head; cur != nil && cur.Next != nil; {
        if cur.Val != cur.Next.Val {
            pre, cur = cur, cur.Next
        } else {
            point := cur
            for ;point != nil && point.Val == cur.Val; point = point.Next{
            }
            if pre != nil {
                pre.Next = point
            } else {
                head = point
            }
            cur = point
        }
    }
    return head
}
