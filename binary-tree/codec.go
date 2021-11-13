package main
// https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree/

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

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	left := this.serialize(root.Left)
	right := this.serialize(root.Right)
	return left + "," + right + "," + strconv.Itoa(root.Val) 
}

// Deserializes your encoded data to tree.
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
		value, _ := strconv.Atoi(last)
		root := &TreeNode{value, nil, nil}
		root.Right = help()
		root.Left = help()
		return root
	}
	return help()
}



