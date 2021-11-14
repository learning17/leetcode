package main
// https://leetcode-cn.com/problems/longest-increasing-subsequence/

import "math"

func lengthOfLIS(nums []int) int {
	size := len(nums)
	if size < 2 {
		return size
	}
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = math.MaxInt64
	}
	var maxSize int
	for i := 0; i < size; i++ {
		pos := searchLeft(arr, nums[i])
		arr[pos] = nums[i]
		if pos+1 > maxSize {
			maxSize = pos + 1
		}
	}
	return maxSize
}

func searchLeft(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

/*func lengthOfLIS(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}
	dp := make([]int, size)
	maxSize := 1
	for i := 0; i < size; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		maxSize = max(maxSize, dp[i])
	}
	return maxSize
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}*/
