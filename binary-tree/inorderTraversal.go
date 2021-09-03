package main

// https://leetcode-cn.com/problems/binary-tree-inorder-traversal/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	stack := []*TreeNode{}
	p := root
	for {
		for {
			if p == nil {
				break
			} 
			stack = append(stack, p)
			p = p.Left
		}
		size := len(stack)
		if size == 0 {
			break
		}
		node := stack[size-1]
		stack = stack[:size-1]
		res = append(res, node.Val)
		p = node.Right
	}
	return res
}
