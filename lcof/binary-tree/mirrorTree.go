package main
// https://leetcode-cn.com/problems/er-cha-shu-de-jing-xiang-lcof/
// 二叉树的镜像

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	mirrorTree(root.Left)
	mirrorTree(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}
