package main
// https://leetcode-cn.com/problems/decode-ways/
import (
	"strconv"
)
func numDecodings(s string) int {
    size := len(s)
    if size == 0 || s[0] == '0' {
        return 0
    }
    dp := make([]int, size+1)
    dp[1], dp[0] = 1, 1
    for i := 2; i < size+1; i++ {
        value,_ := strconv.Atoi(s[i-2:i])
        if value >= 10 && value <= 26 {
            dp[i] = dp[i-2]
        }
        if s[i-1] != '0' {
            dp[i] += dp[i-1]
        }
    }
    return dp[size]
}
