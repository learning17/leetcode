package main
// https://leetcode-cn.com/problems/sorted-matrix-search-lcci/

func searchMatrix(matrix [][]int, target int) bool {
	row_size := len(matrix)
	if row_size == 0 {
		return false
	}
	col_size := len(matrix[0])
	return search(matrix, 0, 0, row_size - 1, col_size - 1, target)
}

func search(matrix [][]int, r1, c1, r2, c2 int, target int) bool {
	if r1 > r2 || c1 > c2 {
		return false
	}
	mid_r, mid_c := r1 + (r2-r1)/2, c1 + (c2-c1)/2
	if matrix[mid_r][mid_c] == target {
		return true
	} else if matrix[mid_r][mid_c] > target {
		return search(matrix, r1, c1, r2, mid_c-1, target) || search(matrix, r1, mid_c, mid_r-1, c2, target)
	} else {
		return search(matrix, mid_r+1, c1, r2, c2, target) || search(matrix, r1, mid_c+1, mid_r, c2, target)
	}
}
