package main
// https://leetcode-cn.com/problems/super-egg-drop/

import (
	"math"
	"fmt"
)

func main() {
	res := superEggDrop(2, 6)
	fmt.Println(res)
}

func superEggDrop(k int, n int) int {
	if k == 1 || n == 0 {
		return n
	}
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, k+1)
		dp[i][1] = i
	}
	for i := 1; i < n+1; i++ {
		for j := 2; j < k+1; j++ {
			dp[i][j] = math.MaxInt64
			for t := 1; t <= i; t++ {
				dp[i][j] = min(dp[i][j], 1+max(dp[t-1][j-1], dp[i-t][j]))
			}
		}
	}
	return dp[n][k]
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
