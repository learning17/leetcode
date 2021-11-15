package main
// https://leetcode-cn.com/problems/trapping-rain-water/

func trap(height []int) int {
	size := len(height)
	if size < 3 {
		return 0
	}
	leftMaxHeight := make([]int, size)
	rightMaxHeight := make([]int, size)
	var leftMax, rightMax int
	for i := 0; i < size; i++ {
		if height[i] > leftMax {
			leftMax = height[i]
		}
		leftMaxHeight[i] = leftMax
		if height[size-i-1] > rightMax {
			rightMax = height[size-i-1]
		}
		rightMaxHeight[size-i-1] = rightMax
	}
	var sum  int
	for i := 0; i < size; i++ {
		sum += min(leftMaxHeight[i], rightMaxHeight[i]) - height[i]
	}
	return sum
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
