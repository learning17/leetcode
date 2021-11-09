package main
// https://www.nowcoder.com/practice/a9fec6c46a684ad5a3abd4e365a9d362

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func threeOrders( root *TreeNode ) [][]int {
	return [][]int{preorderTraversal(root), inorderTraversal(root), postorderTraversal(root)}
}

func preorderTraversal(root *TreeNode) []int {
	var stack []*TreeNode
	var res []int
	for root != nil || len(stack) > 0 {
		for root != nil {
			res = append(res, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return res
}

func inorderTraversal(root *TreeNode) []int {
	var stack []*TreeNode
	var res []int
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return res
}

func postorderTraversal(root *TreeNode) []int {
	var stack []*TreeNode
	var res []int
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
			pre, root = root, nil
		} else {
			stack = append(stack, root)
			root = root.Right
		}
	}
	return res
}
