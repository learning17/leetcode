package main
// https://www.nowcoder.com/practice/9c158345c867466293fc413cff570356

func maxProduct( arr []float64 ) float64 {
	size := len(arr)
	if size == 0 {
		return 0
	}
	var dp [2]float64
	dp[0], dp[1] = arr[0], arr[0]
	maxProduct := arr[0]
	for i := 1; i < size; i++ {
		dp[0], dp[1] = min(min(arr[i], dp[0]*arr[i]), dp[1]*arr[i]), max(max(arr[i], dp[0]*arr[i]), dp[1]*arr[i])
		maxProduct = max(maxProduct, dp[1])
	}
	return maxProduct
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

