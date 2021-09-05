package main
// https://leetcode-cn.com/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof/

func maxSubArray(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}
	dp := make([]int, size)
	dp[0] = nums[0]
	maxSum := nums[0]
	for i := 1; i < size; i++ {
		dp[i] = max(nums[i], dp[i-1] + nums[i])
		maxSum = max(maxSum, dp[i])
	}
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
