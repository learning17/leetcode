package main

// https://leetcode-cn.com/problems/sudoku-solver/

func solveSudoku(board [][]byte)  {
	size := len(board)
	isValid := func(r, c int, b byte) bool {
		for i := 0; i < size; i++ {
			if board[r][i] == b || board[i][c] == b {
				return false
			}
		}
		r_i, c_i := r/3, c/3
		for i := 3*r_i; i < 3*(r_i+1); i++ {
			for j := 3*c_i; j < 3*(c_i+1); j++ {
				if board[i][j] == b {
					return false
				}
			}
		}
		return true
	} 

	arr := []byte{'1','2','3','4','5','6','7','8','9'}
	var backtrack func(int, int) bool 
	backtrack = func(r, c int) bool {
		if r == size {
			return true
		}
		if c == size {
			return backtrack(r+1, 0)
		}
		if board[r][c] != '.' {
			return backtrack(r, c+1)
		}
		for i := 0; i < 9; i++ {
			if !isValid(r, c, arr[i]) {
				continue
			}
			board[r][c] = arr[i]
			if backtrack(r, c+1) {
				return true
			}
			board[r][c] = '.'
		}
		return false
	}
	backtrack(0 ,0 )
}

