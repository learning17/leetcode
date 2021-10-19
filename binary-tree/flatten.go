package main
// https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode)  {
	stack := []*TreeNode{}
	list := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			list = append(list, root)
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	var pre *TreeNode
	for _, node := range list {
		if pre == nil {
			pre = node
			root = node
		} else {
			pre.Right, pre.Left = node, nil
			pre = pre.Right
		}
	}
}
