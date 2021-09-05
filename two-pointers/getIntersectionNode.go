package main

// https://leetcode-cn.com/problems/intersection-of-two-linked-lists-lcci/

type ListNode struct {
	Val int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pA, pB := headA, headB
	flagA, flagB := false, false
	for {
		if pA == nil {
			if !flagA {
				pA, flagA = headB, true
			} else {
				return nil
			}
		}
		if pB == nil {
			if !flagB {
				pB, flagB = headA, true
			} else {
				return nil
			}
		}
		if pA == pB {
			return pA
		}
		pA, pB = pA.Next, pB.Next
	}
}

