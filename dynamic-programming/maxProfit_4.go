package main

// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/

func maxProfit(prices []int) int {
	size := len(prices)
	if size == 0 {
		return 0
	}
	dp := make([][]int, size)
	for i := 0; i < size; i++ {
		dp[i] = make([]int, 3)
	}
	dp[0][0] = -prices[0]
	for i := 1; i < size; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][2] - prices[i])
		dp[i][1] = dp[i-1][0] + prices[i]
		dp[i][2] = max(dp[i-1][2], dp[i-1][1])
	}
	return max(dp[size-1][1], dp[size-1][2])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
