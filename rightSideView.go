package main

// https://leetcode-cn.com/problems/binary-tree-right-side-view/
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	queue := [][]*TreeNode{[]*TreeNode{root}}
	var res []int
	for {
		size := len(queue)
		if size == 0 {
			return res
		}
		nodes := queue[0]
		queue = queue[1:]
		var tmpNodes []*TreeNode
		for i,node := range nodes {
			if i == 0 {
				res = append(res, node.Val)
			}
			if node.Right != nil {
				tmpNodes = append(tmpNodes, node.Right)
			}
			if node.Left != nil {
				tmpNodes = append(tmpNodes, node.Left)
			}
		}
		if len(tmpNodes) != 0 {
			queue = append(queue, tmpNodes)
		}
	}
}
