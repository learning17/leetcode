package main
// https://leetcode-cn.com/problems/arithmetic-slices/

func numberOfArithmeticSlices(nums []int) int {
	size := len(nums)
	if size < 3 {
		return 0
	}
	dp := make([]int, size)
	for i := 2; i < size; i++ {
		if nums[i] - nums[i-1] == nums[i-1] - nums[i-2] {
			dp[i] = dp[i-1] + 1
		}
	}
	sum := 0
	for i := 0; i < size; i++ {
		sum += dp[i]
	}
	return sum
}

