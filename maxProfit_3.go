package main

// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii/

func maxProfit(prices []int) int {
	size := len(prices)
	if size == 0 {
		return 0
	}
	buy1, sell1, buy2, sell2 := -prices[0], 0, -prices[0], 0
	for i := 1; i < size; i++ {
		buy1 = max(buy1, -prices[i])
		sell1 = max(sell1, buy1+prices[i])
		buy2 = max(buy2, sell1 - prices[i])
		sell2 = max(sell2, buy2+prices[i])
	}
	return sell2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
