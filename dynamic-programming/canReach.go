package main
// https://leetcode-cn.com/problems/jump-game-vii/
/*
跳跃游戏 VII
给你一个下标从 0 开始的二进制字符串 s 和两个整数 minJump 和 maxJump
一开始，你在下标 0 处，且该位置的值一定为 '0'
同时满足如下条件时，你可以从下标 i 移动到下标 j 处
i + minJump <= j <= min(i + maxJump, s.length - 1) 且s[j] == '0'
如果你可以到达 s 的下标 s.length - 1 处，请你返回 true ，否则返回 false
*/

func canReach(s string, minJump int, maxJump int) bool {
	size := len(s)
	if size == 0 {
		return true
	}
	if s[size-1] != '0' {
		return false
	}
	dp := make([]bool, size)
	dp[0] = true
	left, right := -1, -1
	for i := 0; i < size; i++ {
		if s[i] == '1' || !dp[i]{
			continue
		}
		tmpLeft, tmpRight := i+minJump, i+maxJump
		if tmpLeft < right {
			left, right = right, tmpRight
		} else {
			left, right = tmpLeft, tmpRight
		}
		for j := left; j < size && j <= right; j++ {
			if s[j] == '0' {
				dp[j] = true
			}
		}
	}
	return dp[size-1]
}

