package main

func maxLength( arr []int ) int {
	dict := make(map[int]bool)
	maxSize := 1
	for i, j := 0, 0; i <= j && j < len(arr); {
		flag, ok := dict[arr[j]]
		if !ok || !flag {
			dict[arr[j]] = true
			j++
		} else {
			dict[arr[i]] = false 
			i++
		}
		if j - i > maxSize {
			maxSize = j - i
		}
	}
	return maxSize
}
