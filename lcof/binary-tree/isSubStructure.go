package main
// https://leetcode-cn.com/problems/shu-de-zi-jie-gou-lcof/
// 树的子结构
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	if check(A, B) {
		return true
	}
	return isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func check(root1, root2 *TreeNode) bool {
	if root2 == nil {
		return true
	}
	if root1 == nil {
		return false
	}
	if root1.Val != root2.Val {
		return false
	}
	return check(root1.Left, root2.Left) && check(root1.Right, root2.Right)
}
