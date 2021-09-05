package main
// https://leetcode-cn.com/problems/house-robber/

func rob(nums []int) int {
	var dp0, dp1 int
	size := len(nums)
	if size == 0 {
		return 0
	}
	if size == 1 {
		return nums[0]
	}
	dp0, dp1 = nums[0], max(nums[0], nums[1])
	for i := 2; i < size;i++ {
		dp0, dp1 = dp1, max(dp1, dp0 + nums[i])
	}
	return dp1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

