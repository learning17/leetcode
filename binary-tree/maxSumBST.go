package main
// https://leetcode-cn.com/problems/maximum-sum-bst-in-binary-tree/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

var maxSum int
func maxSumBST(root *TreeNode) int {
	maxSum = 0
	help(root)
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func help(root *TreeNode) {
	if _, ok := isValidBST(root, nil, nil); ok {
		return
	}
	help(root.Left)
	help(root.Right)
}

func isValidBST(root, min, max *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	if min != nil && root.Val <= min.Val {
		return 0, false
	}
	if max != nil && root.Val >= max.Val {
		return 0, false
	}
	leftValue, leftOk := isValidBST(root.Left, min, root)
	rightValue, rightOk := isValidBST(root.Right, root, max)
	if !leftOk || !rightOk {
		return 0, false
	}
	if maxSum < leftValue+rightValue+root.Val {
		maxSum = leftValue+rightValue+root.Val
	}
	return leftValue+rightValue+root.Val, true
}
