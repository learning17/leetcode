package main
// https://leetcode-cn.com/problems/path-sum-iii/
/*
路径总和 III
给定一个二叉树的根节点 root ，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的 路径 的数目
路径 不需要从根节点开始，也不需要在叶子节点结束
*/
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	left := pathSum(root.Left, targetSum)
	right := pathSum(root.Right, targetSum)
	var num int
	var backtrack func(node *TreeNode, target int)
	backtrack = func(node *TreeNode, target int) {
		if node == nil {
			return
		}
		if target - node.Val == 0 {
			num++
		}
		backtrack(node.Left, target- node.Val)
		backtrack(node.Right, target- node.Val)
	}
	backtrack(root, targetSum)
	return left+right+num
}
