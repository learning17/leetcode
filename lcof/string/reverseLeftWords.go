package main
// https://leetcode-cn.com/problems/zuo-xuan-zhuan-zi-fu-chuan-lcof/
// 左旋转字符串

func reverseLeftWords(s string, n int) string {
	if n >= len(s) {
		return s
	}
	return s[n:] + s[:n]
}
