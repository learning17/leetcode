package main
// https://leetcode-cn.com/problems/0ynMMM/
// 直方图最大矩形面积

func largestRectangleArea(heights []int) int {
	var stack []int
	var maxArea int
	for i := 0; i < len(heights); i++ {
		for len(stack) > 0 && heights[i] < stack[len(stack)-1] {
			curI := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			area := heights[curI] * (i-curI)
			if area > maxArea {
				maxArea = area
			}
		}
		if len(stack) == 0 || stack[len(stack)-1] <= heights[i] {
			stack = append(stack, i)
			continue
		}
	}
	var rightI, minHeight int
	for len(stack) > 0 {
		curI := stack[len(stack)-1]
		if rightI == 0 {
			rightI = curI
		}
		area := (rightI-curI+1) * heights[curI]
		if area > maxArea {
			maxArea = area
		}
		minHeight = heights[curI]
	}
	return maxArea
}
