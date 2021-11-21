package main
// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv/
// 买卖股票的最佳时机 IV

func maxProfit(k int, prices []int) int {
	size := len(prices)
	if size == 0 {
		return 0
	}
	dp := make([][][]int, size)
	for i := 0; i < size; i++ {
		dp[i] = make([][]int, k+1)
		for j := 0; j < k+1; j++ {
			dp[i][j] = make([]int, 2)
			dp[0][j][1] = -prices[0]
		}
	}
	for i := 1; i < size; i++ {
		for j := 1; j < k+1; j++ {
			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1] + prices[i])
			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0] - prices[i])
		}
	}
	var maxProfit int
	for i := 1; i < k+1; i++ {
		if dp[size-1][k][0] > maxProfit {
			maxProfit = dp[size-1][k][0]
		}
	}
	return maxProfit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

