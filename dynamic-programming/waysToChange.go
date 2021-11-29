package main
// https://leetcode-cn.com/problems/coin-lcci/
// 硬币

func waysToChange(n int) int {
	if n <= 1 {
		return 1
	}
	coins := []int{1, 5, 10, 25}
	dp := make([][]int, 5)
	for i := 0; i < 5; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = 1
	}

	for i := 1; i < 5; i++ {
		for j := 1; j < n+1; j++ {
			if j < coins[i-1] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-coins[i]]
			}
		}
	}
	return dp[4][n] % 1000000007
}

