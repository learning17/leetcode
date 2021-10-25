package main
// https://www.nowcoder.com/practice/3afe6fabdb2c46ed98f06cfd9a20f2ce
func findElement( mat [][]int ,  n int ,  m int ,  x int ) []int {
	ans := []int{-1, -1}
	var help func(int, int, int, int) bool
	help = func(rowStart int, rowEnd int, colStart int, colEnd int) bool {
		if rowStart < 0 || rowStart > rowEnd || colStart < 0 || colStart > colEnd {
			return false
		}
		rowMid := rowStart + (rowEnd - rowStart)/2
		colMid := colStart + (colEnd - colStart)/2
		if mat[rowMid][colMid] == x {
			ans[0], ans[1] = rowMid, colMid
			return true
		} else if mat[rowMid][colMid] > x {
			return help(rowStart, rowEnd, colStart, colMid-1) || help(rowStart, rowMid-1, colMid, colEnd)
		} else {
			return help(rowStart, rowEnd, colMid+1, colEnd) || help(rowMid+1, rowEnd, colStart, colMid)
		}
	}
	help(0, n-1, 0, m-1)
	return ans
}
