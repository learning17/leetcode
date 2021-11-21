package main
// https://leetcode-cn.com/problems/ba-shu-zi-fan-yi-cheng-zi-fu-chuan-lcof/
// 把数字翻译成字符串
import (
	"strconv"
)

func translateNum(num int) int {
	s := strconv.Itoa(num)
	size := len(s)
	if size < 2 {
		return size
	}
	dp := make([]int, size+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i < size+1; i++ {
		v,_ := strconv.Atoi(s[i-2:i])
		if v >= 10 && v < 26 {
			dp[i] = dp[i-2]
		} 
		dp[i] += dp[i-1]
	}
	return dp[size]
}
