package main
import "math"
// https://leetcode-cn.com/problems/perfect-squares/

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i < n+1;i++ {
		dp[i] = math.MaxInt64
		for j := 1; j*j <= i; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
