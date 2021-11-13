package main
// https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/
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
	return build(dict, inorder, postorder, 0, len(inorder)-1, 0, len(postorder)-1)
}

func build(dict map[int]int, inorder, postorder []int, iLeft, iRight, pLeft, pRight int) *TreeNode {
	if iLeft > iRight || pLeft > pRight {
		return nil
	}
	root := &TreeNode{postorder[pRight], nil, nil}
	pos := dict[postorder[pRight]]
	leftSize := pos - iLeft
	root.Left = build(dict, inorder, postorder, iLeft, pos-1, pLeft, pLeft+leftSize-1)
	root.Right = build(dict, inorder, postorder, pos+1, iRight, pLeft+leftSize, pRight-1)
	return root
}
