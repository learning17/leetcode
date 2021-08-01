package main

// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/

func maxProfit(prices []int) int {
	size := len(prices)
	if size <= 1 {
		return 0
	}
	dp := make([][]int, size)
	for i := 0; i < size; i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][1] = -prices[0]
	for i := 1; i < size; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0] - prices[i])
	}
	return dp[size-1][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
