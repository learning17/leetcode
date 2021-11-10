package main
// https://www.nowcoder.com/practice/b736e784e3e34731af99065031301bca

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func FindPath( root *TreeNode ,  expectNumber int ) [][]int {
	var res [][]int
	var find func(node *TreeNode, num int, path []int)
	find = func(node *TreeNode, num int, path []int) {
		if node == nil {
			return
		}
		path = append(path, node.Val)
		num -= node.Val
		if node.Left == nil && node.Right == nil && num == 0 {
			res = append(res, append([]int{}, path...))
			return
		}
		find(node.Left, num, path)
		find(node.Right, num, path)
	}
	var path []int
	find(root, expectNumber, path)
	return res
}

