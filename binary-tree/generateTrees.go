package main
// https://leetcode-cn.com/problems/unique-binary-search-trees-ii/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	
	var help func(start, end int) []*TreeNode
	help = func(start, end int) []*TreeNode {
		var res []*TreeNode
		if start > end {
			return res
		}
		if start == end {
			res = append(res, &TreeNode{start, nil, nil})
			return res
		}
		for i := start; i <= end; i++ {
			leftNodes := help(start, i-1)
			rightNodes := help(i+1, end)
			for _,left := range leftNodes {
				for _, right := range rightNodes {
					node := &TreeNode{i, left, right}
					res = append(res, node)
				}
			} 
			if len(leftNodes) == 0 {
				for _, right := range rightNodes {
					node := &TreeNode{i, nil, right}
					res = append(res, node)
				}
			}
			if len(rightNodes) == 0 {
				for _, left := range leftNodes {
					node := &TreeNode{i, left, nil}
					res = append(res, node)
				}
			}
		}
		return res
	}
	return help(1, n)
}
