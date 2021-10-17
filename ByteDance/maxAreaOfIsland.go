package main
// https://leetcode-cn.com/problems/max-area-of-island/

func maxAreaOfIsland(grid [][]int) int {
	row_size := len(grid)
	if row_size == 0 {
		return 0
	}
	col_size := len(grid[0])
	r_dirs := []int{-1, 1, 0, 0}
	c_dirs := []int{0, 0, -1, 1}

	help := func(r, c int) int {
		stack := [][]int{[]int{r,c}}
		grid[r][c] = 0
		area := 1
		for len(stack) > 0 {
			p := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			for i := 0; i < 4; i++ {
				r_n, c_n := p[0] + r_dirs[i], p[1] + c_dirs[i]
				if r_n < 0 || r_n >= row_size || c_n < 0 || c_n >= col_size || grid[r_n][c_n] == 0 {
					continue
				}
				grid[r_n][c_n] = 0
				stack = append(stack, []int{r_n, c_n})
				area++
			}
		}
		return area
	}
	maxArea := 0
	for r := 0; r < row_size; r++ {
		for c := 0; c < col_size; c++ {
			if grid[r][c] == 0 {
				continue
			}
			area := help(r, c)
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}
