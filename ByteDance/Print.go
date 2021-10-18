package main
// https://www.nowcoder.com/practice/91b69814117f4e8097390d107d2efbe0

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func Print( pRoot *TreeNode ) [][]int {
	res := [][]int{}
	if pRoot == nil {
		return res
	}
	queues := [][]*TreeNode{[]*TreeNode{pRoot}}
	index := 0
	for len(queues) > 0 {
		nodes := queues[len(queues)-1]
		queues = queues[:len(queues)-1]
		tmp := []int{}
		tmpNodes := []*TreeNode{}
		for _, node := range nodes {
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				tmpNodes = append(tmpNodes, node.Left)
			}
			if node.Right != nil {
				tmpNodes = append(tmpNodes, node.Right)
			}
		}
		if len(tmp) > 0 {
			if index % 2 == 1 {
				for i, j := 0, len(tmp)-1; i < j; i, j = i+1, j-1 {
					tmp[i],tmp[j] = tmp[j],tmp[i]
				} 
			}
			res = append(res, tmp)
		}
		if len(tmpNodes) > 0 {
			queues = append(queues, tmpNodes)
		}
		index++
	}
	return res
}

