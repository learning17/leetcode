package main
// https://leetcode-cn.com/problems/longest-common-subsequence/

func longestCommonSubsequence(text1 string, text2 string) int {
	size1, size2 := len(text1), len(text2)
	if size1 == 0 || size2 == 0 {
		return 0
	}
	dp := make([][]int, size1+1)
	for i := 0; i < size1+1; i++ {
		dp[i] = make([]int, size2+1)
	}
	for i := 1; i < size1+1; i++ {
		for j := 1; j < size2+1; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[size1][size2]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
