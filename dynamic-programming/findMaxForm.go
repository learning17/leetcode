package main
// https://leetcode-cn.com/problems/ones-and-zeroes/discuss/439318/0-1bei-bao/

import (
	"strings"
)

func findMaxForm(strs []string, m int, n int) int {
	size := len(strs)
	if size == 0 {
		return 0
	}
	dp := make([][][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([][]int, n+1)
		for j := 0; j < n+1; j++ {
			dp[i][j] = make([]int, size+1)
		}
	}
	for k := 1; k < size+1; k++ {
		count0 := strings.Count(strs[k-1], "0")
		count1 := strings.Count(strs[k-1], "1")
		for i := 0; i < m+1; i++ {
			for j := 0; j < n+1; j++ {
				if count0 > i || count1 > j {
					dp[i][j][k] = dp[i][j][k-1]
				} else {
					dp[i][j][k] = max(dp[i][j][k-1], dp[i-count0][j-count1][k-1]+1)
				}
			}
		}
	}
	return dp[m][n][size]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}


