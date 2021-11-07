package main
// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv/
import "math"

func maxProfit(k int, prices []int) int {
	size := len(prices)
	if size < 2 {
		return 0
	}
	buy := make([][]int, size)
	sell := make([][]int, size)
	for i := 0; i < size; i++ {
		buy[i] = make([]int, k+1)
		sell[i] = make([]int, k+1)
	}
	for i := 0; i < k+1; i++ {
		buy[0][i] = math.MinInt64/2
		sell[0][i] = math.MinInt64/2
	}
	buy[0][0] = -prices[0]
	for i := 1; i < size; i++ {
		buy[i][0] = max(buy[i-1][0], -prices[i])
		for j := 1; j < k+1; j++ {
			buy[i][j] = max(buy[i-1][j], sell[i-1][j]-prices[i])
			sell[i][j] = max(sell[i-1][j], buy[i-1][j-1]+prices[i])
		}
	}
	maxValue := sell[size-1][0]
	for _, v := range sell[size-1] {
		if v > maxValue {
			maxValue = v
		}
	}
	return maxValue
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
