package main
// https://leetcode-cn.com/problems/bu-yong-jia-jian-cheng-chu-zuo-jia-fa-lcof/
/*
不用加减乘除做加法
*/

func add(a int, b int) int {
	for b != 0 {
		a, b = a ^ b, (a&b) << 1
	}
	return a
}

