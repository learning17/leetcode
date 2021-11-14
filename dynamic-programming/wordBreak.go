package main
// https://leetcode-cn.com/problems/word-break/

func wordBreak(s string, wordDict []string) bool {
	size := len(s)
	if size == 0 {
		return true
	}
	dict := make(map[string]int)
	for _,word := range wordDict {
		dict[word] = len(word)
	}
	dp := make([]bool, size+1)
	dp[0] = true
	for i := 1; i < size+1; i++ {
		for word, wordSize := range dict {
			if wordSize > i {
				continue
			}
			if word == s[i-wordSize:i] {
				dp[i] = dp[i] || dp[i-wordSize]
			}
		}
	}
	return dp[size]
}
