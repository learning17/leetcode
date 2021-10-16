package main
// https://www.nowcoder.com/practice/253d2c59ec3e4bc68da16833f79a38e4

type ListNode struct{
    Val int
    Next *ListNode
}

func EntryNodeOfLoop(pHead *ListNode) *ListNode{
	if pHead == nil || pHead.Next == nil {
		return nil
	}
	slow, fast := pHead, pHead
	for ;slow != nil && fast != nil && fast.Next != nil; {
		slow, fast = slow.Next, fast.Next.Next
		if slow == fast {
			break
		}
	}
	if slow == nil {
		return nil
	}
	for fast = pHead; slow != nil && fast != nil;{
		if slow == fast {
			return slow
		}
		slow, fast = slow.Next, fast.Next
	}
	return nil
}
