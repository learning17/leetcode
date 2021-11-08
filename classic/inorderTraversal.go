package main

// https://leetcode-cn.com/problems/binary-tree-inorder-traversal/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	res, stack := []int{}, []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		root = node.Right
	}
	return res
}
