package main
// https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
/*
从前序与中序遍历序列构造二叉树
*/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	dict := make(map[int]int)
	for i := 0; i < len(inorder); i++ {
		dict[inorder[i]] = i
	}
	return build(dict, preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1)
}

func build(dict map[int]int, preorder []int, pLeft, pRight int, inorder []int, iLeft, iRight int) *TreeNode {
	if pLeft > pRight || iLeft > iRight {
		return nil
	}
	pos := dict[preorder[pLeft]]
	size := pos - iLeft
	root := &TreeNode{preorder[pLeft], nil, nil}
	root.Left = build(dict, preorder, pLeft+1, pLeft+size, inorder, iLeft, pos-1)
	root.Right = build(dict, preorder, pLeft+size+1, pRight, inorder, pos+1, iRight)
	return root
}

