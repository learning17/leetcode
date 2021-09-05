package main

import "math"
// https://leetcode-cn.com/problems/01-matrix/
// 从左上角到右下角，从右下角到左上角两次搜索 
func updateMatrix(mat [][]int) [][]int {
	row_size := len(mat)
	if row_size == 0 {
		return mat
	}
	col_size := len(mat[0])
	dp := make([][]int, row_size)
	for i := 0; i < row_size; i++{
		dp[i] = make([]int, col_size)
	}
	for i := 0; i < row_size; i++ {
		for j := 0; j < col_size; j++ {
			dp[i][j] = math.MaxInt64 - 1
			if mat[i][j] == 0 {
				dp[i][j] = 0
			} else {
				if i > 0 {
					dp[i][j] = min(dp[i][j], dp[i-1][j]+1)
				}
				if j > 0 {
					dp[i][j] = min(dp[i][j], dp[i][j-1]+1)
				}
			}
		}
	}

	for i := row_size-1; i >= 0; i-- {
		for j := col_size-1; j >= 0; j-- {
			if mat[i][j] == 0 {
				dp[i][j] = 0
			} else {
				if i < row_size - 1 {
					dp[i][j] = min(dp[i][j], dp[i+1][j]+1)
				}
				if j < col_size - 1 {
					dp[i][j] = min(dp[i][j], dp[i][j+1] + 1)
				}
			}
		}
	}
	return dp
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
