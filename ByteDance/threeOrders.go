package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func threeOrders( root *TreeNode ) [][]int {
	return [][]int{preorderTraversal(root), inorderTraversal(root), postorderTraversal(root)}
}

func preorderTraversal(root *TreeNode) []int {
	res, stack := []int{}, []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			res = append(res, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = root.Right

	}
	return res
}

func inorderTraversal(root *TreeNode) []int {
	res, stack := []int{}, []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		res = append(res, root.Val)
		stack = stack[:len(stack)-1]
		root = root.Right
	}
	return res
}

func postorderTraversal(root *TreeNode) []int {
	res, stack := []int{}, []*TreeNode{}
	var pre *TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Right == nil || root.Right == pre {
			res = append(res, root.Val)
			pre = root
			root = nil
		} else {
			stack = append(stack, root)
			root = root.Right
		}
	}
	return res
}
