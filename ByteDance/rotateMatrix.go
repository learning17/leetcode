package main
// https://www.nowcoder.com/practice/2e95333fbdd4451395066957e24909cc

func rotateMatrix( mat [][]int ,  n int ) [][]int {
	row_size, col_size := (n-1)/2, (n-1)/2
	if n % 2 != 0 {
		col_size--
	}
	for i := 0; i <= row_size; i++ {
		for j := 0; j <= col_size;j++ {
			mat[j][n-1-i], mat[n-1-i][n-1-j], mat[n-1-j][i], mat[i][j] = 
			mat[i][j], mat[j][n-1-i], mat[n-1-i][n-1-j], mat[n-1-j][i]
		}
	}
	return mat
}

