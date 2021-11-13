package main
// https://leetcode-cn.com/problems/binary-search-tree-to-greater-sum-tree/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func bstToGst(root *TreeNode) *TreeNode {
	var sum int
	var help func(node *TreeNode)
	help = func(node *TreeNode) {
		if node == nil {
			return
		}
		help(node.Right)
		sum += node.Val
		node.Val = sum
		help(node.Left)
	}
	help(root)
	return root
}
