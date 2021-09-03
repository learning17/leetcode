package main
// https://leetcode-cn.com/problems/binary-tree-right-side-view/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	var res []int
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
		nodeSize := len(nodes)
		var tmpQueue []*TreeNode
		for i := 0; i < nodeSize; i++ {
			if i == nodeSize-1 {
				res = append(res, nodes[i].Val)
			}
			if nodes[i].Left != nil {
				tmpQueue = append(tmpQueue, nodes[i].Left)
			}
			if nodes[i].Right != nil {
				tmpQueue = append(tmpQueue, nodes[i].Right)
			}
		}
		if len(tmpQueue) != 0 {
			queue = append(queue, tmpQueue)
		}
	}
	return res
}
