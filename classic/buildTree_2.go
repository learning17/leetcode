package main
// https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-postorder-traversal/
/*
根据前序和后序遍历构造二叉树
*/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	dict := make(map[int]int)
	for i := 0; i < len(postorder); i++ {
		dict[postorder[i]] = i
	}
	return build(dict, preorder, 0, len(preorder)-1, postorder, 0, len(postorder)-1)
}

func build(dict map[int]int, preorder []int, prLeft, prRight int, postorder []int, poLeft, poRight int) *TreeNode {
	if prLeft > prRight || poLeft > poLeft {
		return nil
	}
	root := &TreeNode{preorder[prLeft], nil, nil}
	if prLeft + 1 > prRight {
		return root
	}
	pos := dict[preorder[prLeft+1]]
	size := pos - poLeft+1
	root.Left = build(dict, preorder, prLeft+1, prLeft+size, postorder, poLeft, pos)
	root.Right = build(dict, preorder, prLeft+size+1, prRight, postorder, pos+1, poRight-1)
	return root
}
