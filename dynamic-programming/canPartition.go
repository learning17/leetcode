package main
// https://leetcode-cn.com/problems/partition-equal-subset-sum/

func canPartition(nums []int) bool {
	size := len(nums)
	if size == 0 {
		return false
	}
	var sum int
	for _, num := range nums {
		sum += num
	}
	if sum % 2 != 0 {
		return false
	}
	half := sum / 2
	dp := make([][]bool, size+1)
	for i := 0; i < size+1; i++ {
		dp[i] = make([]bool, half+1)
		dp[i][0] = true
	}
	for i := 1; i < size+1; i++ {
		for j := 1; j < half+1; j++ {
			if j < nums[i-1] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
			}
		}
	}
	return dp[size][half]
}


