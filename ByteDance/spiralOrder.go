package main
// https://www.nowcoder.com/practice/7edf70f2d29c4b599693dc3aaeea1d31

func spiralOrder( matrix [][]int ) []int {
	row_size := len(matrix)
	if row_size == 0 {
		return []int{}
	}
	col_size := len(matrix[0])
	arr := make([]int, row_size*col_size)
	arr_size := row_size * col_size
	var loop, r, c,i int
	for {
		row_start, row_end := loop, row_size - loop - 1
		col_start, col_end := loop, col_size - loop - 1
		if row_start > row_end || col_start > col_end {
			break
		}
		r, c = row_start, col_start
		for ;c <= col_end && i < arr_size; c, i = c+1, i+1 {
			arr[i] = matrix[r][c]
		}
		r, c = row_start+1, col_end
		for ;r <= row_end && i < arr_size; r, i = r+1, i+1 {
			arr[i] = matrix[r][c]
		}
		r, c = row_end, col_end-1
		for ;c >= col_start && i < arr_size; c, i = c-1, i+1 {
			arr[i] = matrix[r][c]
		}
		r, c = row_end-1, col_start
		for ;r > row_start && i < arr_size; r, i = r-1, i+1 {
			arr[i] = matrix[r][c]
		}
		loop++
	}
	return arr
}

