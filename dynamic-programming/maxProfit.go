package main
import "math"
// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/

func maxProfit(prices []int) int {
	size := len(prices)
	if size == 0 {
		return 0
	}
	maxProfit, minPrice := 0, math.MaxInt64
	for i := 0; i < size; i++ {
		if minPrice > prices[i] {
			minPrice = prices[i]
		} else {
			maxProfit = max(maxProfit, prices[i] - minPrice)
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
