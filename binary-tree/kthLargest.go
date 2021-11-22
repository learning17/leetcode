package main
// https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof/
// 二叉搜索树的第k大节点

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func kthLargest(root *TreeNode, k int) int {
	var stack []*TreeNode
	var i int
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Right
		}
		i++
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if i == k {
			return root.Val
		}
		root = root.Left
	}
	return root.Val
}
