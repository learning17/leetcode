package main
// https://www.nowcoder.com/practice/185a87cd29eb42049132aed873273e83
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func sumNumbers( root *TreeNode ) int {
	var ans int
	var backtrack func(*TreeNode, []int)
	backtrack = func(node *TreeNode, path []int) {
		if node == nil {
			return
		}
		path = append(path, node.Val)
		if node.Left == nil && node.Right == nil {
			var value int
			for i := 0; i < len(path); i++ {
				value = 10*value+path[i]
			}
			ans += value
			return
		}
		backtrack(node.Left, path)
		backtrack(node.Right, path)
	}
	var path []int
	backtrack(root, path)
	return ans
}
