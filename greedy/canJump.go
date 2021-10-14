package main

// https://leetcode-cn.com/problems/jump-game/

func canJump(nums []int) bool {
	maxPos := nums[0]
	for i := 1; i < len(nums); i++ {
		if maxPos < i {
			return false
		}
		maxPos = max(maxPos, nums[i] + i)
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

