package main
// https://leetcode-cn.com/problems/shu-zhi-de-zheng-shu-ci-fang-lcof/
// 数值的整数次方(快速幂)

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		x, n = 1/x, -n
	}
	var res float64
	for res = 1; n > 0; n = n >> 1 {
		if n & 1 == 1 {
			res *= x
		}
		x *= x
	}
	return res
}
