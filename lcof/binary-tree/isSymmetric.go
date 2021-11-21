package main
// https://leetcode-cn.com/problems/dui-cheng-de-er-cha-shu-lcof/
// 对称的二叉树

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return check(root.Left, root.Right)
}

func check(node1, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 == nil || node2 == nil {
		return false
	}
	if node1.Val != node2.Val {
		return false
	}
	return check(node1.Left, node2.Right) && check(node1.Right, node2.Left)
}


