package main

// https://leetcode-cn.com/problems/spiral-matrix/

func spiralOrder(matrix [][]int) []int {
	row_size := len(matrix)
	if row_size == 0 {
		return []int{}
	}
	col_size := len(matrix[0])
	res := make([]int, row_size*col_size)
	left, right, top, bottom := 0, col_size-1, 0, row_size-1
	index := 0

	for ; left <= right && top <= bottom; {
		for c := left; c <= right; c++ {
			res[index] = matrix[top][c]
			index++
		}
		for r := top+1; r <= bottom; r++ {
			res[index] = matrix[r][right]
			index++
		}
		if left < right && top < bottom {
			for c := right-1; c >= left; c-- {
				res[index] = matrix[bottom][c]
				index++
			}
			for r := bottom-1; r > top; r-- {
				res[index] = matrix[r][left]
				index++
			}
		}
		left++
		top++
		right--
		bottom--
	}
	return res
}
