package main
// https://leetcode-cn.com/problems/minimum-path-sum/

func minPathSum(grid [][]int) int {
	row_size := len(grid)
	if row_size == 0 {
		return 0
	}
	col_size := len(grid[0])
	dp := make([][]int, row_size)
	for i := 0; i < row_size; i++ {
		dp[i] = make([]int, col_size)
		if i == 0 {
			dp[i][0] = grid[i][0]
		} else {
			dp[i][0] = grid[i][0] + dp[i-1][0]
		}
	}
	for j := 1; j < col_size; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	for i := 1; i < row_size; i++ {
		for j := 1; j < col_size; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[row_size-1][col_size-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
