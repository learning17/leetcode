package main
// https://www.nowcoder.com/practice/31c1aed01b394f0b8b7734de0324e00f

func maxWater( arr []int ) int64 {
	size := len(arr)
	if size < 3 {
		return 0
	}
	left := make([]int, size)
	right := make([]int, size)
	left[0], right[size-1] = arr[0], arr[size-1]
	for i := 0; i < size; i++ {
		if i > 0 {
			left[i] = max(left[i-1], arr[i])
		}
		if i < size-1 {
			right[size-2-i] = max(right[size-1-i], arr[size-2-i])
		}
	}
	sum := 0
	for i := 0; i < size; i++ {
		sum += min(left[i], right[i]) - arr[i]
	}
	return int64(sum)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
