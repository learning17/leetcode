package main
// https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree/
/*
二叉树的序列化与反序列化
*/

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

type Codec struct {

}

func Constructor() Codec {
	return Codec{}
}

func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	left := this.serialize(root.Left)
	right := this.serialize(root.Right)
	return left + "," + right + "," + strconv.Itoa(root.Val) 
}

func (this *Codec) deserialize(data string) *TreeNode {
	arr := strings.Split(data, ",")
	var help func() *TreeNode
	help = func() *TreeNode {
		if len(arr) == 0 {
			return nil
		}
		last := arr[len(arr)-1]
		arr = arr[:len(arr)-1]
		if last == "#" {
			return nil
		}
		val, _ := strconv.Atoi(last)
		root := &TreeNode{val, nil, nil}
		root.Right = help()
		root.Left = help()
		return root
	}
	return help()
}

