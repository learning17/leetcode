package main
// https://leetcode-cn.com/problems/ones-and-zeroes/discuss/439318/0-1bei-bao/
/*
01 背包（当前物品拿或者不拿，每个物品被限定只能选择0个或1个）：一和零
给你一个二进制字符串数组 strs 和两个整数 m 和 n
请你找出并返回strs的最大子集的长度，该子集中最多有 m 个 0 和 n 个 1
*/

import (
	"strings"
)

func findMaxForm(strs []string, m int, n int) int {
	size := len(strs)
	if size == 0 {
		return 0
	}
	dp := make([][][]int, size+1)
	for i := 0; i < size+1; i++ {
		dp[i] = make([][]int, m+1)
		for j := 0; j < m+1; j++ {
			dp[i][j] = make([]int, n+1)
		}
	}

	for i := 1; i < size+1; i++ {
		num0, num1 := strings.Count(strs[i-1], "0"), strings.Count(strs[i-1], "1") 
		for j := 0; j < m+1; j++ {
			for k := 0; k < n+1; k++ {
				if num0 > j || num1 > k {
					dp[i][j][k] = dp[i-1][j][k]
				} else {
					dp[i][j][k] = max(dp[i-1][j][k], dp[i-1][j-num0][k-num1]+1)
				}
			}
		}
	}
	return dp[size][m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

