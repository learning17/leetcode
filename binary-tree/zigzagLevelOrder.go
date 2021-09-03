package main
// https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	var queue [][]*TreeNode
	queue = append(queue, []*TreeNode{root})
	level := 0
	for {
		size := len(queue)
		if size == 0 {
			break
		}
		nodes := queue[0]
		queue = queue[1:size]
		level %= 2
		var tmpRes []int
		var tmpQueue []*TreeNode
		for i := len(nodes) - 1; i >= 0; i-- {
			tmpRes = append(tmpRes, nodes[i].Val)
			if level == 0 {
				if nodes[i].Left != nil {
					tmpQueue = append(tmpQueue, nodes[i].Left)
				}
				if nodes[i].Right != nil {
					tmpQueue = append(tmpQueue, nodes[i].Right)
				}
			} else {
				if nodes[i].Right != nil {
					tmpQueue = append(tmpQueue, nodes[i].Right)
				}
				if nodes[i].Left != nil {
					tmpQueue = append(tmpQueue, nodes[i].Left)
				}
			}
		}
		level++
		if len(tmpRes) != 0 {
			res = append(res, tmpRes)
		}
		if len(tmpQueue) != 0 {
			queue = append(queue, tmpQueue)
		}
	}
	return res
}
