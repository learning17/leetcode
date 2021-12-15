package main
// https://leetcode-cn.com/problems/partition-equal-subset-sum/
/*
子集背包:分割等和子集
给你一个只包含正整数的非空数组nums 。
请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等
*/

func canPartition(nums []int) bool {
	sum, size := calSum(nums), len(nums)
	if sum % 2 != 0 || size == 0 {
		return false
	}
	half := sum / 2
	dp := make([][]bool, size+1)
	for i := 0; i < size+1; i++ {
		dp[i] = make([]bool, half+1)
		dp[i][0] = true
	}
	for i := 1; i < size+1; i++ {
		for j := 0; j < half+1; j++ {
			if j < nums[i-1] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
			}
		}
	}
	return dp[size][half]
}

func calSum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}


