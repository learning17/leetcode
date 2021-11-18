package main
// https://leetcode-cn.com/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/
// 二维数组中的查找

func findNumberIn2DArray(matrix [][]int, target int) bool {
	row_size := len(matrix)
	if row_size == 0 {
		return false
	}
	col_size := len(matrix[0])
	return find(matrix, 0, row_size-1, 0, col_size-1, target)
}

func find(matrix [][]int, r1, r2, c1, c2 int, target int) bool {
	if r1 > r2 || c1 > c2 {
		return false
	}
	rm, cm := r1 +(r2-r1)/2, c1+(c2-c1)/2
	if matrix[rm][cm] == target {
		return true
	} else if matrix[rm][cm] < target {
		return find(matrix, r1, r2, cm+1, c2, target) || find(matrix, rm+1, r2, c1, cm, target)
	} else {
		return find(matrix, r1, r2, c1, cm-1, target) || find(matrix, r1, rm-1, cm, c2, target)
	}
}
