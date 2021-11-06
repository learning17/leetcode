package main
// https://www.nowcoder.com/practice/3911a20b3f8743058214ceaa099eeb45

func minMoney( arr []int ,  aim int ) int {
	dp := make([]int, aim+1)
	for i := 0; i < aim+1; i++ {
		dp[i] = -1
	}
	dp[0] = 0
	for i := 1; i < aim+1; i++ {
		for _,money := range arr {
			if i-money < 0 || dp[i-money] == -1 {
				continue
			}
			if dp[i] == -1 {
				dp[i] = dp[i-money]+1
				continue
			}
			dp[i] = min(dp[i], dp[i-money]+1)
		}
	}
	return dp[aim]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
