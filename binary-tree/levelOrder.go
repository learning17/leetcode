package main

// https://leetcode-cn.com/problems/binary-tree-level-order-traversal/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	var queue [][]*TreeNode
	queue = append(queue, []*TreeNode{root})
	for {
		size := len(queue)
		if size == 0 {
			break
		}
		nodes := queue[0]
		queue = queue[1:size]
		var tmpRes []int
		var tmpQueue []*TreeNode
		for _, node := range nodes {
			tmpRes = append(tmpRes, node.Val)
			if node.Left != nil {
				tmpQueue = append(tmpQueue, node.Left)
			}
			if node.Right != nil {
				tmpQueue = append(tmpQueue, node.Right)
			}
		}
		if len(tmpQueue) != 0 {
			queue = append(queue, tmpQueue)
		}
		if len(tmpRes) != 0 {
			res = append(res, tmpRes)
		}
	}
	return res
}
