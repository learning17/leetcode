package main
// https://www.nowcoder.com/practice/947f6eb80d944a84850b0538bf0ec3a5

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func Convert( pRootOfTree *TreeNode ) *TreeNode {
	var stack []*TreeNode
	var root, pre *TreeNode = nil, nil
	for pRootOfTree != nil || len(stack) > 0 {
		for pRootOfTree != nil {
			stack = append(stack, pRootOfTree)
			pRootOfTree = pRootOfTree.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		pRootOfTree = node.Right
		if root == nil {
			root = node
		} else {
			pre.Right = node
			node.Left = pre
		}
		pre = node
	}
	return root
}

