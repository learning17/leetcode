package main
// https://www.nowcoder.com/practice/508378c0823c423baa723ce448cbfd0c
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func hasPathSum( root *TreeNode ,  sum int ) bool {
	if root == nil {
		return false
	}
	var help func(node *TreeNode, value int) bool
	help = func(node *TreeNode, value int) bool {
		if node == nil && value == 0 {
			return true
		}
		if value == 0 || node == nil {
			return false
		}
		return help(node.Left, value-node.Val) || help(node.Right, value-node.Val)
	}
	if root.Left == nil && root.Right != nil {
		return help(root.Right, sum - root.Val)
	}
	if root.Right == nil && root.Left != nil {
		return help(root.Left, sum-root.Val)
	}
	return help(root, sum)
}
