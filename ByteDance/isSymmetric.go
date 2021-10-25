package main
// https://www.nowcoder.com/practice/1b0b7f371eae4204bc4a7570c84c2de1

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isSymmetric( root *TreeNode ) bool {
	if root == nil {
		return true
	}
	return help(root.Left, root.Right)
}

func help(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return help(left.Left, right.Right) && help(left.Right, right.Left)
}

