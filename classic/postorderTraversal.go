package main
// https://leetcode-cn.com/problems/binary-tree-postorder-traversal/
/*
二叉树的后序非递归遍历
按照访问左子树——右子树——根节点的方式遍历这棵树
*/
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	res, stack := []int{}, []*TreeNode{}
	var pre *TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Right == nil || root.Right == pre {
			res = append(res, root.Val)
			pre = root
			root = nil
		} else {
			stack = append(stack, root)
			root = root.Right
		}
	}
	return res
}
