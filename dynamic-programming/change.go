package main
// https://leetcode-cn.com/problems/coin-change-2/

func change(amount int, coins []int) int {
	size := len(coins)
	if size == 0 {
		return 0
	}
	dp := make([][]int, size+1)
	for i := 0; i < size+1; i++ {
		dp[i] = make([]int, amount+1)
		dp[i][0] = 1
	}
	dp[0][0] = 0
	for i := 1; i < size+1; i++ {
		for j := 1; j < amount+1; j++ {
			if j < coins[i-1] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]]
			}
		}
	}
	return dp[size][amount]
}
