package main

// https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := [][]*TreeNode{[]*TreeNode{root}}
	level := 0
	res := [][]int{}
	for ; len(queue) > 0;{
		nodes := queue[0]
		queue = queue[1:]
		level %= 2
		tmpRes, tmpNodes := []int{}, []*TreeNode{}
		for i :=  len(nodes)- 1; i >= 0; i-- {
			tmpRes = append(tmpRes, nodes[i].Val)
			if level == 0 {
				if nodes[i].Left != nil {
					tmpNodes = append(tmpNodes, nodes[i].Left)
				}
				if nodes[i].Right != nil {
					tmpNodes = append(tmpNodes, nodes[i].Right)
				}
			} else {
				if nodes[i].Right != nil {
					tmpNodes = append(tmpNodes, nodes[i].Right)
				}
				if nodes[i].Left != nil {
					tmpNodes = append(tmpNodes, nodes[i].Left)
				}
			}
		}
		if len(tmpRes) > 0 {
			res = append(res, tmpRes)
		}
		if len(tmpNodes) > 0 {
			queue = append(queue, tmpNodes)
		}
		level++
	}
	return res
}
