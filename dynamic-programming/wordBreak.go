package main
// https://leetcode-cn.com/problems/word-break/
/*
单词拆分
给你一个字符串 s 和一个字符串列表 wordDict 作为字典，
判定 s 是否可以由空格拆分为一个或多个在字典中出现的单词
*/

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
