package main
// https://www.nowcoder.com/practice/da785ea0f64b442488c125b441a4ba4a

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func maxPathSum( root *TreeNode ) int {
	if root == nil {
		return 0
	}
	maxSum := root.Val
	var help func(node *TreeNode) int
	help = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := max(0, help(node.Left))
		right := max(0, help(node.Right))
		sum := node.Val + left + right
		if sum > maxSum {
			maxSum = sum
		}
		return node.Val + max(left, right)
	}
	help(root)
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
