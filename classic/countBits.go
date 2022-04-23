package main

// https://leetcode-cn.com/problems/w3tCBm/
// 前 n 个数字二进制中 1 的个数
/*
对于正整数 xx，如果可以知道最大的正整数 yy，使得 y \le xy≤x 且 yy 是 22 的整数次幂，则 yy 的二进制表示中只有最高位是 11，其余都是 00，此时称 yy 为 xx 的「最高有效位」。令 z=x-yz=x−y，显然 0 \le z<x0≤z<x，则 \textit{bits}[x]=\textit{bits}[z]+1bits[x]=bits[z]+1。
*/

func countBits(n int) []int {
	bits := make([]int, n+1)
	var highBit int
	for i := 1; i < n+1; i++ {
		if i&(i-1) == 0 {
			highBit = i
		}
		bits[i] = bits[i-highBit] + 1
	}
	return bits
}
