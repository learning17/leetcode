package main

//https://leetcode-cn.com/problems/longest-palindromic-substring/

func longestPalindrome(s string) string {
	size := len(s)
	if size <= 1 {
		return s
	}
	dp := make([][]bool, size)
	for i := 0; i < size; i++ {
		dp[i] = make([]bool, size)
		dp[i][i] = true
	}
	maxSize, left, right := 0, 0, 0
	for l := 1; l < size; l++ {
		for i := 0; i < size; i++ {
			j := i+l
			if j > size - 1 {
				break
			}
			if s[i] == s[j] && (i+1 ==j || dp[i+1][j-1]){
				dp[i][j] = true
			}
			if dp[i][j] && l > maxSize {
				maxSize, left, right = l, i, j
			}
		}
	}
	return s[left:right+1]
}
