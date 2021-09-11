package main
// https://leetcode-cn.com/problems/word-search/

func exist(board [][]byte, word string) bool {
	row_size := len(board)
	if row_size == 0 {
		return false
	}
	col_size := len(board[0])
	visited := make([][]bool, row_size)
	for i := 0; i < row_size; i++ {
		visited[i] = make([]bool, col_size)
	}
	word_size := len(word)
	dire_x, dire_y := []int{0, 0, -1, 1}, []int{-1, 1, 0, 0}
	var backtrack func(i, j, k int) bool 
	backtrack = func(i, j, k int) bool {
		if board[i][j] != word[k] {
			return false
		}
		if k == word_size - 1 {
			return true
		}
		visited[i][j] = true
		for l := 0; l < 4; l++{
			new_x, new_y := i+dire_x[l], j + dire_y[l]
			if new_x < 0 || new_y < 0 || new_x >= row_size || new_y >= col_size || visited[new_x][new_y] {
				continue
			}
			if backtrack(new_x, new_y, k+1) {
				return true
			}
		}
		visited[i][j] = false
		return false
	}
	for i := 0; i < row_size; i++ {
		for j := 0; j < col_size; j++ {
			if backtrack(i, j, 0) {
				return true
			}
		}
	}
	return false
}
