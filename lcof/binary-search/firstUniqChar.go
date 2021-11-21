package main
// https://leetcode-cn.com/problems/di-yi-ge-zhi-chu-xian-yi-ci-de-zi-fu-lcof/
// 第一个只出现一次的字符

func firstUniqChar(s string) byte {
	dict := make(map[byte]int)
	for i := 0; i < len(s);i++ {
		dict[s[i]]++
	}
	for i := 0; i < len(s);i++ {
		if dict[s[i]] == 1 {
			return s[i]
		}
	}
	return ' '
}
