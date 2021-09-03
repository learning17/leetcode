package main
// https://leetcode-cn.com/problems/zhong-jian-er-cha-shu-lcof/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	dict := make(map[int]int)
	for i,v := range inorder {
		dict[v] = i
	}
	return helper(0, len(preorder)-1, preorder, 0, len(inorder)-1, inorder, dict)
}

func helper(preLeft, preRight int, preorder []int, inLeft, inRight int, inorder []int, dict map[int]int) *TreeNode{
	if preLeft > preRight || inLeft > inRight {
		return nil
	}
	rootVal  := preorder[preLeft]
	index := dict[rootVal]
	leftSize := index - inLeft
	left  := helper(preLeft+1, preLeft+leftSize, preorder, inLeft, inLeft+leftSize-1, inorder, dict)
	right := helper(preLeft+1+leftSize, preRight, preorder, index+1, inRight, inorder, dict)
	return &TreeNode{
		Val: rootVal,
		Left: left,
		Right: right,
	}
}

