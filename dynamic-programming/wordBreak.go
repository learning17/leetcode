package main
// https://leetcode-cn.com/problems/word-break/

func wordBreak(s string, wordDict []string) bool {
	size := len(s)
	if size == 0 {
		return true
	}
	dict := make(map[string]int)
	for _, word := range wordDict {
		dict[word] = len(word)
	}
	dp := make([]bool, size+1)
	dp[0] = true

	for i := 1; i < size+1; i++ {
		for _,word := range wordDict {
			if dict[word] > i {
				continue
			}
			if word == s[i-dict[word]:i] {
				dp[i] = dp[i] || dp[i-dict[word]]
			}
		}
	}
	return dp[size]
}
