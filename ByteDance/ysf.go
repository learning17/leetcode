package main
// https://www.nowcoder.com/practice/41c399fdb6004b31a6cbb047c641ed8a
// 环形链表的约瑟夫问题
// https://leetcode-cn.com/problems/yuan-quan-zhong-zui-hou-sheng-xia-de-shu-zi-lcof/solution/jian-zhi-offer-62-yuan-quan-zhong-zui-ho-dcow/

type Node struct {
	Val int
	Next *Node
}

func ysf(n int, m int) int {
	var ans int
	for i := 2; i < n+1; i++ {
		ans = (ans + m) % i
	}
	return ans
}

/*
func ysf( n int ,  m int ) int {
	root := &Node{Val:1}
	node := root
	for i := 2; i <= n; i++ {
		node.Next = &Node{Val:i}
		node = node.Next
	}
	node.Next = root

	pre, start := node, root
	for i := 0; i < n-1; i++ {
		for j := 0; j < m-1; j++ {
			pre, start = pre.Next, start.Next
		}
		pre.Next, start = start.Next, start.Next
	}
	return pre.Val
}
*/

