package main
// https://leetcode-cn.com/problems/implement-rand10-using-rand7/
// 用 Rand7() 实现 Rand10()
func rand10() int {
	res := 41
	for res > 40 {
		res = (rand7() - 1) * 7 + rand7()
	}
	return 1 + (res-1) % 10
}

func rand7() int {
	return 1
}
