package main
// https://leetcode-cn.com/problems/number-of-islands/
// 岛屿数量

func numIslands(grid [][]byte) int {
	row_size := len(grid[0])
	if row_size == 0 {
		return 0
	}
	col_size := len(grid[0])
	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || r >= row_size || c < 0 || c >= col_size || grid[r][c] == '0' {
			return 
		}
		grid[r][c] = '0'
		dfs(r-1, c)
		dfs(r, c-1)
		dfs(r+1, c)
		dfs(r, c+1)
	}
	var num int
	for r := 0; r < row_size; r++ {
		for c := 0; c < col_size; c++ {
			if grid[r][c] == '0' {
				continue
			}
			num++
			dfs(r,c)
		}
	}
	return num
}
