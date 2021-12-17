package main

// 小堆
type SmallHeap []int

func (s *SmallHeap) Push(x int) {
	*s = append(*s, x)
	size := len(*s)
	parent, i := (size-2)/2, size-1
	for parent >= 0 && (*s)[parent] > (*s)[i] {
		(*s)[parent], (*s)[i] = (*s)[i], (*s)[parent]
		parent, i = (parent-1)/2, parent
	}
}

func (s *SmallHeap) Pop() int {
	x := (*s)[0]
	size := len(*s)
	(*s)[0] = (*s)[size-1]
	*s = (*s)[:size-1]
	size--
	for parent, i := 0, 1; i < size; parent, i = i, 2*i+1 {
		if i+1 < size && (*s)[i+1] < (*s)[i] {
			i++
		}
		if (*s)[parent] < (*s)[i] {
			break
		}
		(*s)[parent], (*s)[i] = (*s)[i], (*s)[parent]
	}
	return x
}

// 大堆
type BigHeap []int

func (s *BigHeap) Push(x int) {
	*s = append(*s, x)
	size := len(*s)
	parent, i := (size-2)/2, size-1
	for parent >= 0 && (*s)[parent] < (*s)[i] {
		(*s)[parent], (*s)[i] = (*s)[i], (*s)[parent]
		parent, i = (parent-1)/2, parent
	}
}

func (s *BigHeap) Pop() int {
	x := (*s)[0]
	size := len(*s)
	(*s)[0] = (*s)[size-1]
	*s = (*s)[:size-1]
	size--
	for parent, i := 0, 1; i < size; parent, i = i, 2*i+1 {
		if i+1 < size && (*s)[i+1] > (*s)[i] {
			i++
		}
		if (*s)[parent] > (*s)[i] {
			break
		}
		(*s)[parent], (*s)[i] = (*s)[i], (*s)[parent]
	}
	return x
}
