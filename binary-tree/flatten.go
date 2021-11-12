package main
// https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode)  {
	flattenRecursion(root) // 递归
	flattenNoRecursion(root) // 非递归
}

func flattenRecursion(root *TreeNode) {
	if root == nil {
		return
	}
	flattenRecursion(root.Left)
	flattenRecursion(root.Right)
	left, right := root.Left, root.Right
	root.Left, root.Right = nil, left
	for root.Right != nil {
		root = root.Right
	}
	root.Right = right
}

func flattenNoRecursion(root *TreeNode) {
	if root == nil {
		return
	}
	var stack []*TreeNode
	var list []*TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			list = append(list, root)
			root = root.Left
		}
		root = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	for i := 0; i < len(list); i++ {
		list[i].Left, list[i].Right = nil, nil
		if i > 0 {
			list[i-1].Right = list[i]
		}
	}
}
