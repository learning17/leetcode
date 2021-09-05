package main
// https://leetcode-cn.com/problems/longest-palindromic-subsequence/

func longestPalindromeSubseq(s string) int {
	size := len(s)
	if size == 0 {
		return 0
	}
	dp := make([][]int, size)
	for i := 0; i < size; i++ {
		dp[i] = make([]int, size)
		dp[i][i] = 1
	}
	for l := 1; l < size; l++ {
		for i := 0; i < size; i++ {
			j := i+l
			if j >= size {
				break
			}
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][size-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
