package main

// https://leetcode-cn.com/problems/jump-game-vi/

import (
	"math"
)

func maxResult(nums []int, k int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}
	dp := make([]int, size)
	for i := 0; i < size; i++ {
		dp[i] = math.MinInt64
	}
	dp[0] = nums[0]
	for i := 0; i < size; i++ {
		for j := i+1; j < size && j <= i+k; j++ {
			dp[j] = max(dp[j], dp[i]+nums[j])
			if dp[j] >= dp[i] {
				break
			}
		}
	}
	return dp[size-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
