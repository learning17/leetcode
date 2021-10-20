package main
// https://www.nowcoder.com/practice/c9480213597e45f4807880c763ddd5f0

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func solve( xianxu []int ,  zhongxu []int ) []int {
	root := reConstructBinaryTree(xianxu, zhongxu)
	ans := []int{}
	if root == nil {
		return ans
	}
	stacks := [][]*TreeNode{[]*TreeNode{root}}
	for len(stacks) > 0 {
		nodes := stacks[len(stacks)-1]
		stacks = stacks[:len(stacks)-1]
		if len(nodes) > 0 {
			ans = append(ans, nodes[len(nodes)-1].Val)
		}
		tmpNodes := []*TreeNode{}
		for _, node := range nodes {
			if node.Left != nil {
				tmpNodes = append(tmpNodes, node.Left)
			} 
			if node.Right != nil {
				tmpNodes = append(tmpNodes, node.Right)
			}
		}
		if len(tmpNodes) > 0 {
			stacks = append(stacks, tmpNodes)
		}
	}
	return ans
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

