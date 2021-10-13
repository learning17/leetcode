package main
// https://www.nowcoder.com/practice/04a5560e43e24e9db4595865dc9c63a3

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func levelOrder( root *TreeNode ) [][]int {
	res, stack := [][]int{}, [][]*TreeNode{}
	if root == nil {
		return res
	}
	stack = append(stack, []*TreeNode{root})
	for len(stack) > 0 {
		nodes := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		tmpRes, tmpStack := []int{}, []*TreeNode{}
		for _,node := range nodes {
			tmpRes = append(tmpRes, node.Val)
			if node.Left != nil {
				tmpStack = append(tmpStack, node.Left)
			}
			if node.Right != nil {
				tmpStack = append(tmpStack, node.Right)
			}
		}
		if len(tmpRes) > 0 {
			res = append(res, tmpRes)
		}
		if len(tmpStack) > 0 {
			stack = append(stack, tmpStack)
		}
	}
	return res
}
