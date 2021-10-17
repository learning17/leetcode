package main
// https://www.nowcoder.com/practice/eac1c953170243338f941959146ac4bf

import (
	"sort"
)

func MLS( arr []int ) int {
	size := len(arr)
	if size < 2 {
		return size
	}
	sort.Ints(arr)
	dp := make([]int, size)
	dp[0] = 1
	maxSize := 1
	for i := 1; i < size; i++ {
		if arr[i] == arr[i-1] {
			dp[i] = dp[i-1]
		} else if arr[i] == arr[i-1] + 1 {
			dp[i] = dp[i-1]+1
		} else {
			dp[i] = 1
		}
		if dp[i] > maxSize {
			maxSize = dp[i]
		}
	}
	return dp[size-1]
}
