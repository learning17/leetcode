package main
// https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-postorder-traversal/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	dict := make(map[int]int)
	for i,v := range postorder {
		dict[v] = i
	}
	return helper(0, len(preorder)-1, preorder, 0, len(postorder)-1, postorder, dict)
}

func helper(preLeft, preRight int, preorder []int, postLeft, postRight int, postorder [] int, dict map[int]int)*TreeNode {
	if preLeft > preRight || postLeft > postRight {
		return nil
	}
	if preLeft == preRight {
		return &TreeNode{Val: preorder[preLeft]}
	}
	index := dict[preorder[preLeft+1]]
	leftSize := index - postLeft + 1
	left := helper(preLeft+1, preLeft+leftSize, preorder, postLeft, index, postorder, dict)
	right := helper(preLeft+leftSize+1, preRight, preorder, index+1, postRight, postorder, dict)
	return &TreeNode{
		Val: preorder[preLeft],
		Left: left,
		Right: right,
	}
}
