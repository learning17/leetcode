package main
// https://www.nowcoder.com/practice/4eaccec5ee8f4fe8a4309463b807a542

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isContains( root1 *TreeNode ,  root2 *TreeNode ) bool {
	if root1 == nil && root2 == nil {
		return true
	} else if root1 == nil || root2 == nil {
		return false
	}
	if root1.Val == root2.Val {
		return isContains(root1.Left, root2.Left) && isContains(root1.Right, root2.Right)
	} else {
		return isContains(root1.Left, root2) || isContains(root1.Right, root2)
	}
}

