package main
import "sort"
// https://leetcode-cn.com/problems/russian-doll-envelopes/
/*
俄罗斯套娃信封问题
当另一个信封的宽度和高度都比这个信封大的时候，这个信封就可以放进另一个信封里，如同俄罗斯套娃一样
请计算 最多能有多少个 信封能组成一组“俄罗斯套娃”信封
*/

func maxEnvelopes(envelopes [][]int) int {
	size := len(envelopes)
	if size == 0 {
		return 0
	}
	sort.Slice(envelopes, func (i,j int) bool {return envelopes[i][1] > envelopes[j][1]})
	maxSize := 1
	dp := make([]int, size)
	dp[0] = 1
	for i := 1; i < size; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if envelopes[j][0] > envelopes[i][0] && envelopes[j][1] > envelopes[i][1] {
				dp[i] = max(dp[i], dp[j] + 1)
			}
		}
		maxSize = max(maxSize, dp[i])
	}
	return maxSize
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
