package main

// https://leetcode-cn.com/problems/maximal-square/

func maximalSquare(matrix [][]byte) int {
	row_size := len(matrix)
	if row_size == 0 {
		return 0
	}
	col_size := len(matrix[0])
	dp := make([][]int, row_size)
	maxArea := 0
	for i := 0; i < row_size; i++ {
		dp[i] = make([]int, col_size)
		if matrix[i][0] == '1' {
			dp[i][0] = 1
			maxArea = 1
		}
	}
	for j := 0; j < col_size; j++ {
		if matrix[0][j] == '1' {
			dp[0][j] = 1
			maxArea = 1
		}
	}
	for i := 1; i < row_size; i++ {
		for j := 1; j < col_size; j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
			}
			if dp[i][j] > maxArea {
				maxArea = dp[i][j]
			}
		}
	}
	return maxArea*maxArea
}

func min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b <= a && b <= c {
		return b
	}
	return c
}
