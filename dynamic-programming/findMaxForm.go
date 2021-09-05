package main
// https://leetcode-cn.com/problems/ones-and-zeroes/discuss/439318/0-1bei-bao/

func findMaxForm(strs []string, m int, n int) int {
	size := len(strs)
	if size == 0 {
		return 0
	}
	dp := make([][][]int, size+1)
	for k := 0; k < size+1; k++ {
		dp[k] = make([][]int, m+1)
		for i := 0; i < m+1; i++ {
			dp[k][i] = make([]int, n+1)
		}
	}

	for k := 1; k < size+1; k++ {
		num_0, num_1 := count(strs[k-1])
		for i := 0; i < m+1;i++ {
			for j := 0; j < n + 1; j++ {
				if num_0 > i || num_1 > j {
					dp[k][i][j] = dp[k-1][i][j]
				} else {
					dp[k][i][j] = max(dp[k-1][i][j], dp[k-1][i-num_0][j-num_1])
				}
			}
		}
	}
	return dp[size][m][n]
}

func count(s string)(int, int) {
	var num_0, num_1 int
	for _,c := range s {
		if c == '1' {
			num_1++
		} else {
			num_0++
		}
	}
	return num_0, num_1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
