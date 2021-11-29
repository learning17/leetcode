package main
// https://leetcode-cn.com/problems/spiral-matrix/
// 螺旋矩阵

func spiralOrder(matrix [][]int) []int {
	var res []int
	row_size := len(matrix)
	if row_size == 0 {
		return res 
	}
	col_size := len(matrix[0])
	rowTop, rowBottom := 0, row_size-1 
	colLeft, colRight := 0, col_size-1

	for rowTop <= rowBottom && colLeft <= colRight {
		tmp := GetRowTop(matrix, colLeft, colRight, rowTop)
		res  = append(res, tmp...)
		rowTop++
		tmp = GetColRight(matrix, rowTop, rowBottom, colRight)
		res = append(res, tmp...)
		colRight--
		tmp = GetRowBottom(matrix, colLeft, colRight, rowBottom)
		res = append(res, tmp...)
		rowBottom--
		tmp = GetColLeft(matrix, rowTop, rowBottom, colLeft)
		res = append(res, tmp...)
		colLeft++
	}
	return res[:row_size*col_size]
}

func GetRowTop(matrix [][]int, cStart, cEnd, r int) []int {
	var res []int
	for c := cStart; c <= cEnd; c++ {
		res = append(res, matrix[r][c])
	}
	return res
}

func GetRowBottom(matrix [][]int, cStart, cEnd, r int) []int {
	var res []int
	for c := cEnd; c >= cStart; c-- {
		res = append(res, matrix[r][c])
	}
	return res
}

func GetColLeft(matrix [][]int, rStart, rEnd, c int) []int {
	var res []int
	for r := rEnd; r >= rStart; r-- {
		res = append(res, matrix[r][c])
	}
	return res
}

func GetColRight(matrix [][]int, rStart, rEnd, c int) []int {
	var res []int
	for r := rStart; r <= rEnd; r++ {
		res = append(res, matrix[r][c])
	}
	return res
}
