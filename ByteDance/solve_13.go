package main
// https://www.nowcoder.com/practice/046a55e6cd274cffb88fc32dba695668
import (
	"strconv"
)
func solve( nums string ) int {
	size := len(nums)
	if size == 0 || nums[0] == '0' {
		return 0
	}
	dp := make([]int, size)
	dp[0] = 1
	for i := 1; i < size; i++ {
		if nums[i] != '0' {
			dp[i] = dp[i-1]
		}
		v,_ := strconv.Atoi(nums[i-1:i+1])
		if v >= 10 && v <= 26 {
			if i == 1 {
				dp[i]++
			} else {
				dp[i] += dp[i-2]
			}
		}
	}
	return dp[size-1]
}

