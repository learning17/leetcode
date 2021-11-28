package main
// https://leetcode-cn.com/problems/top-k-frequent-elements/
// 前 K 个高频元素
import (
	"container/heap"
)

func topKFrequent(nums []int, k int) []int {
	dict := make(map[int]int)
	for _, num := range nums {
		dict[num]++
	}
	b := BigHeap{}
	heap.Init(&b)
	var i int
	for num, count := range dict {
		if i < k {
			heap.Push(&b, []int{num, count})
			continue
		}
		if b[0][1] > count {
			continue
		}
		heap.Pop(&b)
		heap.Push(&b, []int{num, count})
	}
	var ans []int
	for i := 0; i < len(b); i++ {
		ans = append(ans, b[i][0])
	}
	return ans
}

type BigHeap [][]int

func (b BigHeap) Len() int {return len(b)}
func (b BigHeap) Less(i, j int) bool {return b[i][1] < b[j][1]}
func (b BigHeap) Swap(i, j int) {b[i], b[j] = b[j], b[i]}

func (b *BigHeap) Push(x interface{}) {
	*b = append(*b, x.([]int))
}

func (b *BigHeap) Pop() interface{} {
	x := (*b)[len(*b)-1]
	*b = (*b)[:len(*b)-1]
	return x
}
