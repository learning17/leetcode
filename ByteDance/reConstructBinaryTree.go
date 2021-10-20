package main
// https://www.nowcoder.com/practice/8a19cbe657394eeaac2f6ea9b0f6fcf6

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func reConstructBinaryTree( pre []int ,  vin []int ) *TreeNode {
	dict := make(map[int]int)
	for i, v := range vin {
		dict[v] = i
	}
	return help(pre, vin, dict, 0, len(pre)-1, 0, len(vin)-1)
}

func help(pre []int, vin []int, dict map[int]int, start1, end1, start2, end2 int) *TreeNode {
	if start1 > end1 {
		return nil
	}
	root := &TreeNode{pre[start1], nil, nil}
	pos := dict[pre[start1]]
	size := pos-start2
	root.Left  = help(pre, vin, dict, start1+1, start1+size, start2, pos-1)
	root.Right = help(pre, vin, dict, start1+size+1, end1, pos+1, end2)
	return root
}

