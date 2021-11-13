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
	if !checkLeft(root.Left, root.Val) || !checkRight(root.Right, root.Val) {
		return false
	}
	if !isValidBST(root.Left) || !isValidBST(root.Right) {
		return false
	}
	return true
}

func checkLeft(node *TreeNode, val int) bool {
	if node == nil {
		return true
	}
	if node.Val >= val {
		return false
	}
	return checkLeft(node.Right, val) && checkLeft(node.Left, val)
}

func checkRight(node *TreeNode, val int) bool {
	if node == nil {
		return true
	}
	if node.Val <= val {
		return false
	}
	return checkRight(node.Right, val) && checkRight(node.Left, val)
}
