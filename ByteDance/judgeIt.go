package main
// https://www.nowcoder.com/practice/f31fc6d3caf24e7f8b4deb5cd9b5fa97

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func judgeIt( root *TreeNode ) []bool {
	return []bool{isValidBST(root), isCompletedTree(root)}
}

// 完全二叉树
func isCompletedTree(root *TreeNode) bool {
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		if node == nil {
			break
		}
		queue = queue[1:]
		queue = append(queue, node.Left)
		queue = append(queue, node.Right)
	}
	for len(queue) > 0 {
		if queue[0] != nil {
			return false
		}
		queue = queue[1:]
	}
	return true
}

// 搜索二叉树
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if !less(root.Val, root.Right) || !more(root.Val, root.Left) {
		return false
	}
	return isValidBST(root.Left) && isValidBST(root.Right)
}

func less(value int, node *TreeNode) bool {
	if node == nil {
		return true
	}
	if value >= node.Val {
		return false
	}
	return less(value, node.Left) && less(value, node.Right)
}

func more(value int, node *TreeNode) bool {
	if node == nil {
		return true
	}
	if value <= node.Val {
		return false
	}
	return more(value, node.Left) && more(value, node.Right)
}

