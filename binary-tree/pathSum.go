package main
// https://leetcode-cn.com/problems/path-sum-iii/
// 路径总和 III
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
		backtrack(node.Left, target - node.Val)
		backtrack(node.Right, target - node.Val)
	}
	backtrack(root, targetSum)
	return num + left + right
}
