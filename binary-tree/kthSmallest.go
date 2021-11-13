package main
// https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	arr := kthSmallestNoRecursion(root)
	return arr[k-1]
}

func kthSmallestNoRecursion(root *TreeNode)(res []int) {
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return res
}
