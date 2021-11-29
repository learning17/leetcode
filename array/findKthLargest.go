package main
// https://leetcode-cn.com/problems/kth-largest-element-in-an-array/
// 数组中的第K个最大元素

func findKthLargest(nums []int, k int) int {
	var heap SmallHeap
	for i, num := range nums {
		if i < k {
			heap.Push(num)
		} else if heap[0] < num {
			heap.Pop()
			heap.Push(num)
		}
	}
	return heap[0]
}

type SmallHeap []int

func (s *SmallHeap) Push(x int) {
	*s = append(*s, x)
	size := len(*s)
	i, parent := size - 1, (size - 2)/2
	for parent >= 0 && (*s)[parent] > (*s)[i] {
		(*s)[parent], (*s)[i] = (*s)[i], (*s)[parent]
		i, parent = parent, (parent-1)/2
	}
}

func (s *SmallHeap) Pop() int {
	x := (*s)[0]
	size := len(*s)
	(*s)[0] = (*s)[size-1]
	*s = (*s)[:size-1]
	size--
	for parent, i := 0, 1; i <= size-1; parent, i = i, 2*i+1 {
		if i + 1 < size && (*s)[i+1] < (*s)[i] {
			i++
		}
		if (*s)[parent] < (*s)[i] {
			break
		}
		(*s)[parent], (*s)[i] = (*s)[i], (*s)[parent]
	} 
	return x
}

