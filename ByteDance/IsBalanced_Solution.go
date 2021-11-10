package main
// https://www.nowcoder.com/practice/8b3b95850edb4115918ecebdf1b4d222

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func IsBalanced_Solution( pRoot *TreeNode ) bool {
	_, flag := check(pRoot)
	return flag
}

func check(node *TreeNode)(int, bool) {
	if node == nil {
		return 0, true
	}
	if node.Left == nil && node.Right == nil {
		return 1, true
	}
	leftHeight, leftFlag := check(node.Left)
	rightHeight, rightFlag := check(node.Right)
	if !leftFlag || !rightFlag || abs(leftHeight - rightHeight) > 1 {
		return 0, false
	}
	return max(leftHeight, rightHeight)+1, true 
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

