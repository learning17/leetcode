package main
// https://leetcode-cn.com/problems/bu-yong-jia-jian-cheng-chu-zuo-jia-fa-lcof/

func add(a int, b int) int {
	n, c := a ^ b, (a&b) << 1
	for c != 0 {
		n, c = n ^ c, (n & c) << 1
	}
	return n
}

