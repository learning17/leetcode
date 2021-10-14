package main
// https://www.nowcoder.com/practice/459bd355da1549fa8a49e350bf3df484

func FindGreatestSumOfSubArray( array []int ) int {
	dp := make([]int, len(array))
	dp[0] = array[0]
	maxSum := array[0]
	for i := 1; i < len(array); i++ {
		dp[i] = max(dp[i-1]+array[i], array[i])
		maxSum = max(maxSum, dp[i])
	}
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
