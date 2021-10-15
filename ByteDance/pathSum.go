package main
// https://www.nowcoder.com/practice/840dd2dc4fbd4b2199cd48f2dadf930a

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func pathSum( root *TreeNode ,  sum int ) [][]int {
	var res [][]int
	var backtrack func(*TreeNode, int, []int)
	backtrack = func(node *TreeNode, value int, path []int) {
		if node == nil {
			return
		}
		path = append(path, node.Val)
		value -= node.Val
		if node.Left == nil && node.Right == nil && value == 0 {
			res = append(res, append([]int{}, path...))
			return
		}
		backtrack(node.Left, value, path)
		backtrack(node.Right, value, path)
	}
	var path []int
	backtrack(root, sum, path)
	return res
}
