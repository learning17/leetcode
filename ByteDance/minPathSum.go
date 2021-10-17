package main

// https://www.nowcoder.com/practice/7d21b6be4c6b429bb92d219341c4f8bb

func minPathSum( matrix [][]int ) int {
	row_size := len(matrix)
	if row_size == 0 {
		return 0
	}
	col_size := len(matrix[0])
	for r := 1; r < row_size; r++ {
		matrix[r][0] += matrix[r-1][0]
	}
	for c := 1; c < col_size; c++ {
		matrix[0][c] += matrix[0][c-1]
	}
	for r := 1; r < row_size; r++ {
		for c := 1; c < col_size; c++ {
			matrix[r][c] += min(matrix[r-1][c], matrix[r][c-1])
		}
	}
	return matrix[row_size-1][col_size-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

