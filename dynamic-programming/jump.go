package main

// https://leetcode-cn.com/problems/jump-game-ii/
/*
跳跃游戏 II
最初位于数组的第一个位置
数组中的每个元素代表你在该位置可以跳跃的最大长度
目标是使用最少的跳跃次数到达数组的最后一个位置
*/
import (
	"math"
)
func jump(nums []int) int {
    size := len(nums)
    if size < 2 {
        return 0
    }
    dp := make([]int, size)
    for i := 1; i < size; i++ {
        dp[i] = math.MaxInt64
    }
    for i := 1; i < size; i++ {
        for j := 0; j < i; j++ {
            if nums[j] >= i - j {
                dp[i] = min(dp[i], dp[j]+1)
            }
        }
    }
    return dp[size-1]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
