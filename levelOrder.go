package main
// https://leetcode-cn.com/problems/binary-tree-level-order-traversal/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	queue := [][]*TreeNode{[]*TreeNode{root}}
	for {
		size := len(queue)
		if size == 0 {
			break
		}
		nodes := queue[0]
		queue = queue[1:size]
		var tmpNodes []*TreeNode
		var tmpRes []int
		for _,node := range nodes {
			tmpRes = append(tmpRes, node.Val)
			if node.Left != nil {
				tmpNodes = append(tmpNodes, node.Left)
			}
			if node.Right != nil {
				tmpNodes = append(tmpNodes, node.Right)
			}
		}
		res = append(res, tmpRes)
		if len(tmpNodes) != 0 {
			queue = append(queue, tmpNodes)
		}
	}
	return res
}
