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
	var arr []string
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			arr = append(arr, "#")
			continue
		}
		arr = append(arr, strconv.Itoa(node.Val))
		queue = append(queue, node.Left)
		queue = append(queue, node.Right)
	}
	return strings.Join(arr, ",")
}

func Deserialize( s string ) *TreeNode {
	if len(s) == 0 {
		return nil
	}
	buildNode := func(c string) *TreeNode {
		value, _ := strconv.Atoi(c)
		return &TreeNode{value, nil, nil}
	}
	arr := strings.Split(s, ",")
	root := buildNode(arr[0]) 
	queue := []*TreeNode{root}
	var pos int
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		pos++
		if arr[pos] != "#" {
			node.Left = buildNode(arr[pos])
			queue = append(queue, node.Left)
		}
		pos++
		if arr[pos] != "#" {
			node.Right = buildNode(arr[pos])
			queue = append(queue, node.Right)
		}
	}
	return root
}

