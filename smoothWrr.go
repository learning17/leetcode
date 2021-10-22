package main
// 平滑负载均衡
import (
	"fmt"
)

func main() {
	nodes := []*Node{
		&Node{"a", 5, 5},
		&Node{"b", 1, 1},
		&Node{"c", 1, 1},
	}
	for i := 0; i < 7; i++ {
		node := SmoothWrr(nodes)
		fmt.Println(node.Name)
	}
}
type Node struct{
	Name string
	Current int
	Weight int
}

func SmoothWrr(nodes []*Node) (best *Node) {
	size := len(nodes)
	if size == 0 {
		return nil
	}
	bestNode := nodes[0]
	sumWeight := nodes[0].Current
	for i := 1; i < size; i++ {
		if nodes[i].Current > bestNode.Current {
			bestNode = nodes[i]
		}
		sumWeight += nodes[i].Current
	}
	bestNode.Current -= sumWeight
	for i := 0; i < size; i++ {
		nodes[i].Current += nodes[i].Weight
	}
	return bestNode
}
