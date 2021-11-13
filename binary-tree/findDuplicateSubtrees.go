package main
// https://leetcode-cn.com/problems/find-duplicate-subtrees/

import (
	"strconv"
)
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	var res []*TreeNode
	dict := make(map[string]int)
	var help func (*TreeNode) string
	help = func(node *TreeNode) string {
		if node == nil {
			return "#"
		}
		nodeStr := help(node.Left) + "," + help(node.Right) + "," + strconv.Itoa(node.Val)
		dict[nodeStr]++
		if dict[nodeStr] == 2 {
			res = append(res, node)
		}
		return nodeStr
	}
	help(root)
	return res
}


