package main

func lengthOfLIS(nums []int) int {
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
}
