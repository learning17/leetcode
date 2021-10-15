package main
// https://leetcode-cn.com/problems/jump-game-vii/

func canReach(s string, minJump int, maxJump int) bool {
	size := len(s)
	if size == 0 {
		return false
	}
	if s[size-1] != '0' {
		return false
	}
	dp := make([]bool, size)
	dp[0] = true
	left, right := -1, -1
	for i := 0; i < size; i++ {
		if s[i] == '1' || !dp[i] {
			continue
		}
		tmpLeft, tmpRight := i+minJump, i+maxJump
		if tmpLeft < right {
			left, right = right, tmpRight
		} else {
			left, right = tmpLeft, tmpRight
		}
		flag := false
		for j := left; j <= right && j < size; j++ {
			if s[j] == '0' {
				dp[j] = true
				flag = true
			}
		}
		if !flag {
			return false
		}
	}
	return dp[size-1]
}
