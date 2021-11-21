package main
// https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-lcof/
// 从上到下打印二叉树

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := [][]*TreeNode{[]*TreeNode{root}}
	var res [][]int
	var seq int
	for len(queue) > 0 {
		nodes := queue[0]
		queue = queue[1:]
		var tmpQueue []*TreeNode
		var tmpRes []int
		for _,node := range nodes {
			if seq % 2 == 0 {
				tmpRes = append(tmpRes, node.Val)
			} else {
				tmpRes = append([]int{node.Val}, tmpRes...)
			}
			if node.Left != nil {
				tmpQueue = append(tmpQueue, node.Left)
			}
			if node.Right != nil {
				tmpQueue = append(tmpQueue, node.Right)
			}
		}
		if len(tmpQueue) > 0 {
			queue = append(queue, tmpQueue)
		}
		if len(tmpRes) > 0 {
			res = append(res, tmpRes)
		}
		seq++
	}
	return res
}

