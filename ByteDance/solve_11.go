package main
// https://www.nowcoder.com/practice/0058c4092cec44c2975e38223f10470e

func solve( matrix [][]byte ) int {
	row_size := len(matrix)
	if row_size == 0 {
		return 0
	}
	col_size := len(matrix[0])
	maxSize := 0
	dp := make([][]int, row_size)
	for i := 0; i < row_size; i++ {
		dp[i] = make([]int, col_size)
		if matrix[i][0] == '1' {
			dp[i][0] = 1
			maxSize = 1
		}
	}

	for i := 0; i < col_size; i++ {
		if matrix[0][i] == '1' {
			dp[0][i] = 1
			maxSize = 1
		}
	}
	for i := 1; i < row_size; i++ {
		for j := 1; j < col_size; j++ {
			if matrix[i][j] == '0' {
				continue
			}
			dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1])+1
			if dp[i][j] > maxSize {
				maxSize = dp[i][j]
			}
		}
	}
	return maxSize * maxSize
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
