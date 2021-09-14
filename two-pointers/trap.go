package main

// https://leetcode-cn.com/problems/trapping-rain-water/

func trap(height []int) int {
	size := len(height)
	if size == 0 {
		return 0
	}
	leftMax  := make([]int, size)
	rightMax := make([]int, size)
	leftMax[0] = height[0]
	for i := 1; i < size; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}
	rightMax[size-1] = height[size-1]
	for i := size-1; i >=0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}
	ans := 0
	for i := 0; i < size; i++ {
		ans += min(leftMax[i], rightMax[i]) - height[i]
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
