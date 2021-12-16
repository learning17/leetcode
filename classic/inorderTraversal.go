package main

// https://leetcode-cn.com/problems/binary-tree-inorder-traversal/
/*
二叉树的中序非递归遍历
按照访问左子树——根节点——右子树的方式遍历这棵树
*/
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	res, stack := []int{}, []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return res
}
