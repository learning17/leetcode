package main
// https://www.nowcoder.com/practice/cf7e25aa97c04cc1a68c8f040e71fb84
import (
	"strconv"
	"strings"
)
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func Serialize( root *TreeNode ) string {
	if root == nil {
		return ""
	}
	queue := [][]*TreeNode{[]*TreeNode{root}}
	var arr []string
	for len(queue) > 0 {
		nodes := queue[0]
		queue = queue[1:]
		var tmpNodes []*TreeNode
		for _,node := range nodes {
			if node == nil {
				arr = append(arr, "#")
				continue
			}
			arr = append(arr, strconv.Itoa(node.Val))
			tmpNodes = append(tmpNodes, node.Left)
			tmpNodes = append(tmpNodes, node.Right)
		}
		if len(tmpNodes) > 0 {
			queue = append(queue, tmpNodes)
		}
	}
	return strings.Join(arr, ",")
}

func Deserialize( s string ) *TreeNode {
	if len(s) == 0 {
		return nil
	}
	arr := strings.Split(s, ",")
	val, _ := strconv.Atoi(arr[0])
	root := &TreeNode{Val: val}
	queue := []*TreeNode{root}
	size := len(arr)
	for i := 1; len(queue) > 0 && i < size; i = i+2 {
		node := queue[0]
		queue = queue[1:]
		if arr[i] != "#" {
			val, _ := strconv.Atoi(arr[i])
			node.Left = &TreeNode{Val: val}
			queue = append(queue, node.Left)
		}
		if i + 1 < size && arr[i+1] != "#" {
			val, _ := strconv.Atoi(arr[i+1])
			node.Right = &TreeNode{Val: val}
			queue = append(queue, node.Right)
		}
	}
	return root
}

