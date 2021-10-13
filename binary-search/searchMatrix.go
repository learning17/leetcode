package main

// https://leetcode-cn.com/problems/sorted-matrix-search-lcci/

func searchMatrix(matrix [][]int, target int) bool {
	row_size := len(matrix)
	if row_size == 0 {
		return false
	}
	col_size := len(matrix[0])
	return search(matrix, 0, 0, row_size-1, col_size-1, target)
}

func search(matrix [][]int, r1, c1, r2, c2, target int) bool {
	if r1 > r2 || c1 > c2 {
		return false
	}
	r_m, c_m := r1+(r2-r1)/2, c1+(c2-c1)/2
	if matrix[r_m][c_m] == target {
		return true
	} else if matrix[r_m][c_m] > target {
		return search(matrix, r1, c1, r2, c_m-1, target) || search(matrix, r1, c_m, r_m-1, c2, target)
	} else {
		return search(matrix, r1, c_m+1, r2, c2, target) || search(matrix, r_m+1, c1, r2, c_m, target)
	}
}
