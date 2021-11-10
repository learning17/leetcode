package main
// https://www.nowcoder.com/practice/ff05d44dfdb04e1d83bdbdab320efbcb

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isSymmetrical( pRoot *TreeNode ) bool {
	if pRoot == nil {
		return true
	}
	return check(pRoot.Left, pRoot.Right)
}

func check(node1, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 == nil || node2 == nil || node1.Val != node2.Val {
		return false
	}
	return check(node1.Left, node2.Right) && check(node1.Right, node2.Left)
}
