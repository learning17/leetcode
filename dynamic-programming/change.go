package main
// https://leetcode-cn.com/problems/coin-change-2/
/*
完全背包:零钱兑换 II
给你一个整数数组 coins 表示不同面额的硬币，另给一个整数 amount 表示总金额
请你计算并返回可以凑成总金额的硬币组合数。如果任何硬币组合都无法凑出总金额，返回 0
假设每一种面额的硬币有无限个
*/

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
