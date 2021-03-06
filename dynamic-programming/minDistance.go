package main

// https://leetcode-cn.com/problems/edit-distance/
/* 编辑距离
给你两个单词 word1 和 word2，请你计算出将 word1 转换成 word2 所使用的最少操作数
dp[i][j] = dp[i-1][j-1] if word1[i] == word2[j]
dp[i][j] = min(dp[i][j-1], dp[i-1][j], dp[i-1][j-1]) + 1
*/

func minDistance(word1 string, word2 string) int {
	size1, size2 := len(word1), len(word2)
	if size1 == 0 || size2 == 0 {
		return size1 + size2
	}
	dp := make([][]int, size1+1)
	for i := 0; i < size1+1; i++ {
		dp[i] = make([]int, size2+1)
		dp[i][0] = i
	}
	for i := 0; i < size2+1; i++ {
		dp[0][i] = i
	}
	for i := 1; i < size1+1; i++ {
		for j := 1; j < size2+1; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i][j-1], dp[i-1][j], dp[i-1][j-1]) + 1
			}
		}
	}
	return dp[size1][size2]
}

func min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b <= a && b <= c {
		return b
	}
	return c
}
