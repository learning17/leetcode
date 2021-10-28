package main
// https://www.nowcoder.com/practice/a77b4f3d84bf4a7891519ffee9376df3

import (
	"math"
)

type Interval struct {
	Start int
	End int
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func solve( n int ,  Tree_edge []*Interval ,  Edge_value []int ) int {
	dict := make(map[int]*TreeNode)
	dict[0] = &TreeNode{}
	for i := 0; i < n-1; i++ {
		dict[Tree_edge[i].End] = &TreeNode{Val: Edge_value[i]} 
		if dict[Tree_edge[i].Start].Left == nil {
			dict[Tree_edge[i].Start].Left = dict[Tree_edge[i].End]
		} else {
			dict[Tree_edge[i].Start].Right = dict[Tree_edge[i].End]
		}
	}
	var help func(*TreeNode) int
	maxSum := math.MinInt64
	help = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := max(0, help(node.Left))
		right := max(0, help(node.Right))
		maxSum = max(maxSum, left + right)
		return node.Val + max(left, right) 
	}
	help(dict[0])
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

