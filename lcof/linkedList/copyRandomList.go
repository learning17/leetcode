package main
// https://leetcode-cn.com/problems/fu-za-lian-biao-de-fu-zhi-lcof/
// 复杂链表的复制

type Node struct {
	Val int
	Next *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	dict := make(map[*Node]*Node)
	var deepcopy func(node *Node) *Node
	deepcopy = func(node *Node) *Node {
		if node == nil {
			return nil
		}
		if tmp, ok := dict[node]; ok {
			return tmp
		}
		dict[node] = &Node{node.Val, nil, nil}
		dict[node].Next = deepcopy(node.Next)
		dict[node].Random = deepcopy(node.Random)
		return dict[node]
	}
	return deepcopy(head)
}


