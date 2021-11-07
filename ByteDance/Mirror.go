package main
// https://www.nowcoder.com/practice/a9d0ecbacef9410ca97463e4a5c83be7

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func Mirror( pRoot *TreeNode ) *TreeNode {
	if pRoot == nil || (pRoot.Left == nil && pRoot.Right == nil) {
		return pRoot
	}
	left := Mirror(pRoot.Left)
	right := Mirror(pRoot.Right)
	pRoot.Left, pRoot.Right = right, left
	return pRoot
}

