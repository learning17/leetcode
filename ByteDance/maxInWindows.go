package main
// https://www.nowcoder.com/practice/1624bc35a45c42c0bc17d17fa0cba788

func maxInWindows( num []int ,  size int ) []int {
	if len(num) == 0 || size > len(num) || size == 0 {
		return []int{}
	}
	h := Heap{}
	for i := 0; i < size; i++ {
		h.Insert(num[i])
	}
	ans := []int{h.arr[0]}
	for i := size; i < len(num);i++ {
		j := i-size
		if num[i] != num[j] {
			h.Remove(num[j])
			h.Insert(num[i])
		}
		ans = append(ans, h.arr[0])
	}
	return ans
}

type Heap struct {
	arr []int
	size int
}

func (h *Heap) Insert(value int) {
	h.arr = append(h.arr, value)
	h.size++
	for i := h.size-1; i > 0; {
		parent := (i-1)/2
		if h.arr[parent] >= h.arr[i] {
			break
		}
		h.arr[parent], h.arr[i] = h.arr[i], h.arr[parent]
		parent, i = (parent-1)/2, parent
	}
}

func (h *Heap) Remove(value int) {
	pos := 0
	for ; pos < h.size; pos++ {
		if h.arr[pos] == value {
			break
		}
	}
	if pos == h.size {
		return
	}
	h.arr[pos] = h.arr[h.size-1]
	h.arr = h.arr[:h.size-1]
	h.size--
	h.adjust(pos, h.size)
}

func (h *Heap) adjust(k int, size int) {
	for i := 2*k+1; i < size; i, k = 2*i+1, i {
		if i + 1 < size && h.arr[i+1] > h.arr[i] {
			i++
		}
		if h.arr[i] < h.arr[k] {
			break
		}
		h.arr[i], h.arr[k] = h.arr[k], h.arr[i]
	}
}

