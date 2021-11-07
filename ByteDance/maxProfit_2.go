package main
// https://www.nowcoder.com/practice/9e5e3c2603064829b0a0bbfca10594e9

func maxProfit( prices []int ) int {
	size := len(prices)
	if size == 0 {
		return 0
	}
	dp0, dp1 := 0, -prices[0]
	for i := 1; i < size; i++ {
		dp0, dp1 = max(dp0, dp1+prices[i]), max(dp1, dp0-prices[i])
	}
	return dp0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

