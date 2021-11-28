package main
// https://leetcode-cn.com/problems/w3tCBm/
// 前 n 个数字二进制中 1 的个数

func countBits(n int) []int {
	ans := make([]int, n+1)
	for i := 2; i <= n; i++ {
		ans[i] = ans[i >> 1] + i & 1
	}
	return ans
}
