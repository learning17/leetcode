package main
// https://leetcode-cn.com/problems/jump-game-iii/

func canReach(arr []int, start int) bool {
	size := len(arr)
	visited := make(map[int]bool)
	var help func(int) bool
	help = func(i int) bool {
		if arr[i] == 0 {
			return true
		}
		visited[i] = true
		left, right := i-arr[i], i+arr[i]
		if left >= 0 && !visited[left] && help(left) {
			return true
		}
		if right < size && !visited[right] && help(right) {
			return true
		}
		return false
	}

	return help(start)
}
