package main
// https://leetcode-cn.com/problems/delete-node-in-a-bst/
/*
删除二叉搜索树中的节点
*/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else {
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		minNode := findMinNode(root.Right) 
		root.Val = minNode.Val
		root.Right = deleteNode(root.Right, minNode.Val)
	}
	return root
}

func findMinNode(node *TreeNode) *TreeNode {
	for node.Left != nil {
		node = node.Left
	}
	return node
}
