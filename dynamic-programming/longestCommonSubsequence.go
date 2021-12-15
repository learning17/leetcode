package main
// https://leetcode-cn.com/problems/longest-common-subsequence/
/*
最长公共子序列
给定两个字符串 text1 和 text2，返回这两个字符串的最长 公共子序列 的长度
*/

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
