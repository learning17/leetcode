package main
// https://leetcode-cn.com/problems/sum-of-two-integers/description/

func getSum(a int, b int) int {
	for b != 0 {
		a, b = a^b, (a&b) << 1
	}
	return a
}

