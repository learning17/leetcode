package main
// https://leetcode-cn.com/problems/contiguous-sequence-lcci/

func maxSubArray(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}
	dp, maxValue := nums[0], nums[0]
	for i := 1; i < size; i++ {
		dp = max(nums[i], dp+nums[i])
		maxValue = max(dp, maxValue)
	}
	return maxValue
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
