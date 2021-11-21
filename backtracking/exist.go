package main
// https://leetcode-cn.com/problems/ju-zhen-zhong-de-lu-jing-lcof/
// 矩阵中的路径

func exist(board [][]byte, word string) bool {
	rowSize := len(board)
	if rowSize == 0 {
		return false
	}
	colSize := len(board[0])
	wordSize := len(word)
	visited := make([][]bool, rowSize)
	for i := 0; i < rowSize; i++ {
		visited[i] = make([]bool, colSize)
	}
	rDirs := []int{0, 0, -1, 1}
	cDirs := []int{-1, 1, 0, 0}
	var backtrack func(r, c, i int) bool
	backtrack = func(r, c, i int) bool {
		if board[r][c] != word[i] {
			return false
		}
		if i == wordSize - 1 {
			return true
		}
		visited[r][c] = true
		for k := 0; k < 4; k++ {
			rn, cn := r+rDirs[k], c+cDirs[k]
			if rn < 0 || rn >= rowSize || cn < 0 || cn >= colSize || visited[rn][cn] {
				continue
			}
			if backtrack(rn, cn, i+1) {
				return true
			}
		}
		visited[r][c] = false
		return false
	}
	for r := 0; r < rowSize; r++ {
		for c := 0; c < colSize; c++ {
			if backtrack(r, c, 0) {
				return true
			}
		}
	}
	return false
}
