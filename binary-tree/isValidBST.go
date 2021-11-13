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

func checkValid(root, min, max *TreeNode) bool {
	if root == nil {
		return true
	}
	if min != nil && root.Val <= min.Val {
		return false
	}
	if max != nil && root.Val >= max.Val {
		return false
	}
	return checkValid(root.Left, min, root) && checkValid(root.Right, root, max)
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
