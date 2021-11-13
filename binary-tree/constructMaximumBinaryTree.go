package main
// https://leetcode-cn.com/problems/maximum-binary-tree/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	return build(nums, 0, len(nums)-1)
}

func build(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	pos, max := left, nums[left]
	for i := left+1; i <= right; i++ {
		if nums[i] > max {
			pos, max = i, nums[i]
		}
	}
	root := &TreeNode{max, nil, nil}
	root.Left = build(nums, left, pos-1)
	root.Right = build(nums, pos+1, right)
	return root
}
