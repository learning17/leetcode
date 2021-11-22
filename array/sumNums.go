package main
// https://leetcode-cn.com/problems/qiu-12n-lcof/
// 求1+2+…+n

func sumNums(n int) int {
	var res int
	var sum func(int) bool
	sum = func(n int) bool {
		res += n
		return n > 0 && sum(n-1)
	}
	sum(n)
	return res
}
