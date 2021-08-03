package main

// https://leetcode-cn.com/problems/maximum-subarray/

func maxSubArray(nums []int) int {
	maxSum,tmpSum := nums[0], 0
	for _, num := range nums {
		tmpSum += num
		if tmpSum < num {
			tmpSum = num
		}
		maxSum = max(maxSum, tmpSum)
	}
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

