package main

// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/

func maxProfit(prices []int) int {
	size := len(prices)
	if size == 0 {
		return 0
	}
	minPrice, maxProfit := prices[0], 0
	for i := 0; i < size; i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
		tmp := prices[i] - minPrice
		if maxProfit < tmp {
			maxProfit = tmp
		}
	}
	return maxProfit
}
