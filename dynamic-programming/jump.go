package main

// https://leetcode-cn.com/problems/jump-game-ii/

func jump(nums []int) int {
	size := len(nums)
	if size == 0 {
		return -1
	}
	dp := make([]int, size)
	for i := 0; i < size; i++ {
		dp[i] = size
	}
	dp[0] = 0
	for i := 1; i < size; i++ {
		for j := 0; j < i; j++ {
			if nums[j] < i-j {
				continue
			}
			dp[i] = min(dp[i], dp[j]+1)
		}
	}
	if dp[size-1] == size {
		return -1
	}
	return dp[size-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
