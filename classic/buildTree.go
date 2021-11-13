package main
// https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/

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
	return build(dict, preorder, inorder, 0, len(preorder)-1, 0, len(inorder)-1)
}

func build(dict map[int] int, preorder []int, inorder[] int, pLeft,pRight int, iLeft,iRight int) *TreeNode {
	if pLeft > pRight || iLeft > iRight {
		return nil
	}
	root := &TreeNode{preorder[pLeft], nil, nil}
	pos := dict[preorder[pLeft]]
	leftSize := pos - iLeft
	root.Left = build(dict, preorder, inorder, pLeft+1, pLeft+leftSize, iLeft, pos-1)
	root.Right = build(dict, preorder, inorder, pLeft+leftSize+1, pRight, pos+1, iRight)
	return root
}
