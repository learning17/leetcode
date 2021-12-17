package main
// https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/
/*
从中序与后序遍历序列构造二叉树
*/
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	dict := make(map[int]int)
	for i := 0; i < len(inorder); i++ {
		dict[inorder[i]] = i
	}
	return build(dict, inorder, 0, len(inorder)-1, postorder, 0, len(postorder)-1)
}

func build(dict map[int]int, inorder []int, iLeft, iRight int, postorder []int, pLeft, pRight int) *TreeNode {
	if iLeft > iRight || pLeft > pRight {
		return nil
	}
	pos := dict[postorder[pRight]]
	size := pos - iLeft
	root := &TreeNode{postorder[pRight], nil, nil}
	root.Left = build(dict, inorder, iLeft, pos-1, postorder, pLeft, pLeft+size-1)
	root.Right = build(dict, inorder, pos+1, iRight, postorder, pLeft+size, pRight-1)
	return root
}

