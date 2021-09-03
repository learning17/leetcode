package main
// https://leetcode-cn.com/problems/binary-tree-preorder-traversal/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	stack := []*TreeNode{}
	p := root
	for {
		for{
			if p == nil {
				break
			}
			res = append(res, p.Val)
			stack = append(stack, p)
			p = p.Left
		}
		size := len(stack)
		if size == 0 {
			break
		}
		node := stack[size-1]
		stack = stack[:size-1]
		p = node.Right
	}
	return res
}
