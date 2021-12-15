package main

// https://leetcode-cn.com/problems/jump-game-vi/
/*
跳跃游戏 VI
开始你在下标 0 处
最多可以往前跳 k 步
目标是到达数组最后一个位置,得分 为经过的所有数字之和,你能得到的 最大得分
*/
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
