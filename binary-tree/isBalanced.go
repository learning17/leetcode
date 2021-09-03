package main
// https://leetcode-cn.com/problems/balanced-binary-tree/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	return helper(root) != -1
}

func helper(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left  := helper(root.Left)
	right := helper(root.Right)
	if left == -1 || right == -1 || abs(left - right) > 1 {
		return -1
	}
	if left > right {
		return left+1
	}
	return right+1
}

func abs(value int) int {
	if value < 0 {
		value = -value
	}
	return value
}
