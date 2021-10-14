package main

// https://www.nowcoder.com/practice/e0cc33a83afe4530bcec46eba3325116

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor( root *TreeNode ,  o1 int ,  o2 int ) int {
	node := help(root, o1, o2)
	return node.Val
}

func help(root *TreeNode, o1, o2 int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == o1 || root.Val == o2 {
		return root
	}
	left := help(root.Left, o1, o2)
	right := help(root.Right, o1, o2)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right
}
