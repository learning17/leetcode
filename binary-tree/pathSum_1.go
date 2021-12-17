package main
// https://leetcode-cn.com/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof/
/*
二叉树中和为某一值的路径
给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径
*/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, target int) [][]int {
	var res [][]int
	
	var backtrack func(node *TreeNode, tmpTarget int, path []int)
	backtrack = func(node *TreeNode, tmpTarget int, path []int) {
		if node == nil {
			return
		}
		tmpTarget -= node.Val
		path = append(path, node.Val)
		if tmpTarget == 0 && node.Left == nil && node.Right == nil {
			res = append(res, append([]int{}, path...))
			return
		}
		backtrack(node.Left, tmpTarget, path)
		backtrack(node.Right, tmpTarget, path)
	}
	var path []int
	backtrack(root, target, path)
	return res
}
