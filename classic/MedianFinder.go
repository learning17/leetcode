package main
// https://leetcode-cn.com/problems/shu-ju-liu-zhong-de-zhong-wei-shu-lcof/
// 数据流中的中位数
import (
	"container/heap"
)

type BigHeap []int

func (b BigHeap) Len() int {return len(b)}
func (b BigHeap) Less(i, j int) bool {return b[i] > b[j]}
func (b BigHeap) Swap(i, j int) {b[i], b[j] = b[j], b[i]}

func (b *BigHeap) Push(x interface{}) {
	*b = append(*b, x.(int))
}

func (b *BigHeap) Pop() interface{} {
	x := (*b)[len(*b)-1]
	*b = (*b)[:len(*b)-1]
	return x
}

type SmallHeap []int
func (s SmallHeap) Len() int {return len(s)}
func (s SmallHeap) Less(i, j int) bool {return s[i] < s[j]}
func (s SmallHeap) Swap(i, j int) {s[i], s[j] = s[j], s[i]}

func (s *SmallHeap) Push(x interface{}) {
	*s = append(*s, x.(int))
}

func (s *SmallHeap) Pop() interface{} {
	x := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return x
}

type MedianFinder struct {
	big *BigHeap
	small *SmallHeap
}

func Constructor() MedianFinder {
	m := MedianFinder{big: &BigHeap{}, small: &SmallHeap{}}
	heap.Init(m.big)
	heap.Init(m.small)
	return m
}

func (this *MedianFinder) AddNum(num int)  {
	if this.small.Len() == this.big.Len() {
		heap.Push(this.small, num)
		x := heap.Pop(this.small)
		heap.Push(this.big, x)
	} else {
		heap.Push(this.big, num)
		x := heap.Pop(this.big)
		heap.Push(this.small, x)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.big.Len() == this.small.Len() {
		return float64((*this.big)[0] + (*this.small)[0])/2
	}
	return float64((*this.big)[0])
}
