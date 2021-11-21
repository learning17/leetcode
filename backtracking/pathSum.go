package main
// https://leetcode-cn.com/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof/
// 二叉树中和为某一值的路径
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, target int) [][]int {
	var res [][]int
	var backtrack func(node *TreeNode, sum int, path []int)
	backtrack = func(node *TreeNode, sum int, path []int) {
		if node == nil {
			return
		}
		sum -= node.Val
		path = append(path, node.Val) 
		if node.Left == nil && node.Right == nil && sum == 0 {
			res = append(res, append([]int{}, path...))
			return
		}
		backtrack(node.Left, sum, path)
		backtrack(node.Right, sum, path)
	}
	var path []int
	backtrack(root, target, path)
	return res
}
