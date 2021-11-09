package main
// https://www.nowcoder.com/practice/0bf071c135e64ee2a027783b80bf781d

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func inorderTraversal( root *TreeNode ) []int {
	var stack []*TreeNode
	var res []int
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

