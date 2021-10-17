package main

// https://www.nowcoder.com/practice/0c9664d1554e466aa107d899418e814e

func solve( grid [][]byte ) int {
	r_dirs := []int{-1, 1, 0, 0}
	c_dirs := []int{0, 0, -1, 1}
	row_size := len(grid)
	if row_size == 0 {
		return row_size
	}
	col_size := len(grid[0])
	var help func(int, int)
	help = func(r, c int) {
		grid[r][c] = '0'
		stack := [][]int{[]int{r,c}}
		for len(stack) > 0 {
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			for i := 0; i < 4; i++ {
				r_n, c_n := node[0] + r_dirs[i], node[1] + c_dirs[i]
				if r_n < 0 || r_n >= row_size || c_n < 0 || c_n >= col_size || grid[r_n][c_n] == '0'{
					continue
				}
				grid[r_n][c_n] = '0'
				stack = append(stack, []int{r_n, c_n})
			}
		}
	}
	num := 0
	for r := 0; r < row_size; r++ {
		for c := 0; c < col_size; c++ {
			if grid[r][c] == '0' {
				continue
			}
			num++
			help(r, c)
		}
	}
	return num
}

