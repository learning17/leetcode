package main
// https://leetcode-cn.com/problems/decode-ways/
/*
解码方法
一条包含字母 A-Z 的消息通过以下映射进行了 编码
'A' -> 1
'B' -> 2
...
'Z' -> 26
给你一个只含数字的 非空 字符串 s ，请计算并返回 解码 方法的 总数
*/
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
