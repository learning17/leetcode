package main
// https://leetcode-cn.com/problems/zui-chang-bu-han-zhong-fu-zi-fu-de-zi-zi-fu-chuan-lcof/
// 最长不含重复字符的子字符串

func lengthOfLongestSubstring(s string) int {
	size := len(s)
	if size < 2 {
		return size
	}
	dict := make(map[byte]int)
	var left, right, maxSize int
	for left <= right && right < size {
		for left <= right {
			if v,ok := dict[s[right]]; !ok || v == 0 {
				break
			}
			dict[s[left]]--
			left++
		}
		if v,ok := dict[s[right]]; !ok || v == 0 {
			dict[s[right]]++
			if right - left + 1 > maxSize {
				maxSize = right - left + 1
			}
			right++
		}
	}
	return maxSize
}
