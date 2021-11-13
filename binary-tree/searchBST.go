package main
// https://leetcode-cn.com/problems/search-in-a-binary-search-tree/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	} else if root.Val > val {
		return searchBST(root.Left, val)
	}
	return searchBST(root.Right, val)
}
