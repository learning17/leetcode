package main
// https://www.nowcoder.com/practice/4edf6e6d01554870a12f218c94e8a299

import (
	"sort"
	"container/heap"
)

func minmumNumberOfHost( n int ,  startEnd [][]int ) int {
	sort.Slice(startEnd, func(i, j int) bool {
		if startEnd[i][0] < startEnd[j][0] {
			return true
		}
		if startEnd[i][0] == startEnd[j][0] && startEnd[i][1] < startEnd[j][1] {
			return true
		}
		return false
	})
	size := len(startEnd)
	if size < 2 {
		return size
	}
	h := &EndHeap{startEnd[0][1]}
	heap.Init(h)
	maxNum := 1
	for i := 1; i < len(startEnd); i++ {
		for h.Len() > 0 && startEnd[i][0] >= (*h)[0] {
			heap.Pop(h)
		}
		heap.Push(h, startEnd[i][1])
		if h.Len() > maxNum {
			maxNum = h.Len()
		}
	}
	return maxNum
}

type EndHeap []int

func (e EndHeap) Len() int {return len(e)}
func (e EndHeap) Less(i, j int) bool {return e[i] < e[j]}
func (e EndHeap) Swap(i, j int) {e[i], e[j] = e[j], e[i]}


func (e *EndHeap) Push(x interface{}) {
	*e = append(*e, x.(int))
}

func (e *EndHeap) Pop() interface{} {
	x := (*e)[len(*e)-1]
	*e = (*e)[:len(*e)-1]
	return x
}
