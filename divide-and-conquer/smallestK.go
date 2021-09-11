package main

// https://leetcode-cn.com/problems/smallest-k-lcci/

func smallestK(arr []int, k int) []int {
	size := len(arr)
	if size == 0 {
		return arr
	}
	for i := (k-2)/2; i >= 0; i-- {
		adjust(arr, k, i)
	}
	for i := k; i < size; i++ {
		if arr[0] > arr[i] {
			arr[0] = arr[i]
			adjust(arr, k, 0)
		}
	}
	return arr[:k]
}

func adjust(arr []int, size, i int) {
	for k := 2*i+1; k < size; k=2*k+1 {
		if k + 1 < size && arr[k] < arr[k+1] {
			k++
		}
		if arr[i] < arr[k] {
			arr[i], arr[k] = arr[k], arr[i]
			i = k
		} else {
			break
		}
	}
}

