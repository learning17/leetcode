package main
// https://leetcode-cn.com/problems/insert-into-a-binary-search-tree/
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{val, nil, nil}
	}
	if root.Val < val {
		root.Right = insertIntoBST(root.Right, val)
	}
	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	}
	return root
}

