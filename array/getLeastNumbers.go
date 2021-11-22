package main
// https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof/
// 最小的k个数
import (
	"container/heap"
)
func getLeastNumbers(arr []int, k int) []int {
	h := &Heap{}
	heap.Init(h)
	for i := 0; i < len(arr); i++ {
		heap.Push(h, arr[i])
	}
	var res []int
	for i := 0; i < k; i++ {
		x := heap.Pop(h)
		res = append(res, x.(int))
	}
	return res
}

type Heap []int

func (h Heap) Len() int {return len(h)}
func (h Heap) Less(i, j int) bool {return h[i] < h[j]}
func (h Heap) Swap(i ,j int) {h[i], h[j] = h[j], h[i]}

func (h *Heap) Push(x interface{}){
	*h = append(*h, x.(int))
}

func (h *Heap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}
