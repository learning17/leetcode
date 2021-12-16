package main
// https://leetcode-cn.com/problems/binary-tree-preorder-traversal/
/*
二叉树的前序非递归遍历
按照访问根节点——左子树——右子树的方式遍历这棵树
*/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	res, stack := []int{}, []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			res = append(res, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = root.Right
	}
	return res
}
