package main
// https://leetcode-cn.com/problems/ping-heng-er-cha-shu-lcof/
// 平衡二叉树

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if abs(height(root.Left) - height(root.Right)) > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := height(root.Left)
	right := height(root.Right)
	return max(left, right) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
