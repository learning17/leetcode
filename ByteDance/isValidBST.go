package main
// https://leetcode-cn.com/problems/validate-binary-search-tree/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return true
	}
	if root.Left != nil && root.Val <= root.Left.Val {
		return false
	}
	if root.Right != nil && root.Val >= root.Right.Val {
		return false
	}
	if !more(root.Val, root.Left) || !less(root.Val, root.Right) {
		return false
	}
	return isValidBST(root.Left) && isValidBST(root.Right)
}

func less(value int, node *TreeNode) bool {
	if node == nil {
		return true
	}
	if value >= node.Val {
		return false
	}
	return less(value, node.Left) && less(value, node.Right)
}

func more(value int, node *TreeNode) bool {
	if node == nil {
		return true
	}
	if value <= node.Val {
		return false
	}
	return more(value, node.Left) && more(value, node.Right)
}
