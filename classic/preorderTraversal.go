package main
// https://leetcode-cn.com/problems/binary-tree-preorder-traversal/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	res, stack := []int{}, []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			res = append(res, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return res
}
