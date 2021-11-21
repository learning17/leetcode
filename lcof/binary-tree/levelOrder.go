package main
// https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-lcof/
// 从上到下打印二叉树

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	queue := [][]*TreeNode{[]*TreeNode{root}}
	var res []int
	for len(queue) > 0 {
		nodes := queue[0]
		queue = queue[1:]
		var tmpQueue []*TreeNode
		for _,node := range nodes {
			res = append(res, node.Val)
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
	}
	return res
}

