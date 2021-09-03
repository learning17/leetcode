package main
// https://leetcode-cn.com/problems/binary-tree-postorder-traversal/
// 根——右——左，然后反转得到 左——右——根
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	res := []int{}
	stack := []*TreeNode{}
	p := root
	for {
		for {
			if p == nil {
				break
			}
			res = append(res, p.Val)
			stack = append(stack, p)
			p = p.Right
		}
		size := len(stack)
		if size == 0 {
			break
		}
		node := stack[size-1]
		stack = stack[:size-1]
		p = node.Left
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j -1 {
		res[i], res[j] = res[j],res[i]
	}
	return res
}
