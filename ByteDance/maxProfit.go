package main
// https://www.nowcoder.com/practice/64b4262d4e6d4f6181cd45446a5821ec


func maxProfit( prices []int ) int {
	if len(prices) == 0 {
		return 0
	}
	maxProfit, minPrice := 0, prices[0]
	for i := 1; i < len(prices); i++ {
		minPrice = min(minPrice, prices[i])
		maxProfit = max(maxProfit, prices[i] - minPrice)
	}
	return maxProfit
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
