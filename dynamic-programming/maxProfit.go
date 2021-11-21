package main
import "math"
// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/
// 买卖股票的最佳时机


func maxProfit(prices []int) int {
	var minPrice, maxProfit int
	minPrice = math.MaxInt64
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
		if prices[i] - minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}
	}
	return maxProfit
}

