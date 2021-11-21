package main
// https://leetcode-cn.com/problems/li-wu-de-zui-da-jie-zhi-lcof/
// 礼物的最大价值
func maxValue(grid [][]int) int {
	row_size := len(grid)
	if row_size == 0 {
		return 0
	}
	col_size := len(grid[0])
	dp := make([][]int, row_size)
	for i := 0; i < row_size; i++ {
		dp[i] = make([]int, col_size)
		dp[i][0] = grid[i][0]
		if i > 0 {
			dp[i][0] += dp[i-1][0]
		}
	}
	for i := 0; i < col_size; i++ {
		dp[0][i] = grid[0][i]
		if i > 0 {
			dp[0][i] += dp[0][i-1]
		}
	}

	for r := 1; r < row_size; r++ {
		for c := 1; c < col_size; c++ {
			dp[r][c] = max(dp[r-1][c], dp[r][c-1]) + grid[r][c]
		}
	}
	return dp[row_size-1][col_size-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
