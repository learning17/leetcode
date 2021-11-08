package main
// https://www.nowcoder.com/practice/7a71a88cdf294ce6bdf54c899be967a2

func solve( matrix [][]int ) int {
    row_size := len(matrix)
    if row_size == 0 {
        return 0
    }
    col_size := len(matrix[0])
    r_dirs := []int{-1, 1, 0, 0}
    c_dirs := []int{0, 0, -1, 1}
    maxSize := 1
    var backtrack func(r, c, size int)
    backtrack = func(r, c, size int) {
        if maxSize < size {
            maxSize = size
        }
        for i := 0; i < 4; i++ {
            r_n, c_n := r+r_dirs[i], c+c_dirs[i]
            if r_n < 0 || r_n >= row_size || c_n < 0 || c_n >= col_size || matrix[r_n][c_n] <= matrix[r][c] {
                continue
            }
            backtrack(r_n, c_n, size+1)
        }
    }
    for r := 0; r < row_size; r++ {
        for c := 0; c < col_size; c++ {
            size := 1
            backtrack(r, c, size)
        }
    }
    return maxSize
}
