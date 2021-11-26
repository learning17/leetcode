package main
// https://leetcode-cn.com/problems/iIQa4I/
// 每日温度

func dailyTemperatures(temperatures []int) []int {
	var stack []int
	res := make([]int, len(temperatures))
	for i := 0; i < len(temperatures); i++ {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			cur := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[cur] = i - cur
		}
		stack = append(stack, i)
	}
	return res
}
