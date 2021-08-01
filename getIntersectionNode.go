package main

// https://leetcode-cn.com/problems/intersection-of-two-linked-lists/

type ListNode struct {
	Val int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	pA, pB := headA, headB
	flagA, flagB := false, false
	for {
		if pA == pB {
			return pA
		}
		if pA == nil && !flagA {
			pA, flagA = headB, true
		}
		if pB == nil && !flagB {
			pB, flagB = headA, true
		}
		pA, pB = pA.Next, pB.Next
	}
}

