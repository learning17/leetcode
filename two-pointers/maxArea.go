package main
// https://leetcode-cn.com/problems/container-with-most-water/

func maxArea(height []int) int {
	size := len(height)
	if size == 0 {
		return 0
	}
	maxArea := (size-1)*min(height[0], height[size-1])
	for l, r := 0, size-1; l < r; {
		var tmpArea int
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
		tmpArea = (r - l) * min(height[l], height[r])
		if tmpArea > maxArea {
			maxArea = tmpArea
		}
	}
	return maxArea
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
