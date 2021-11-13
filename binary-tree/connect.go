package main
// https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node/

type Node struct {
	Val int
	Left *Node
	Right *Node
	Next *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	connectRecursion(root.Left, root.Right)
	return root
}

func connectRecursion(node1, node2 *Node) {
	if node1 == nil {
		return
	}
	node1.Next = node2
	connectRecursion(node1.Left, node1.Right)
	connectRecursion(node2.Left, node2.Right)
	connectRecursion(node1.Right, node2.Left)
}
