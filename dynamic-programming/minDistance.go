package main

// https://leetcode-cn.com/problems/edit-distance/

func minDistance(word1 string, word2 string) int {
	size1, size2 := len(word1), len(word2)
	dp := make([][]int, size1+1)
	for i := 0; i < size1+1; i++ {
		dp[i] = make([]int, size2+1)
		dp[i][0] = i
	}
	for j := 0; j < size2+1; j++ {
		dp[0][j] = j
	}

	for i := 1; i < size1+1; i++ {
		for j := 1; j < size2+1; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1])
			} else {
				dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
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
