package main
// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/

func lengthOfLongestSubstring(s string) int {
	size := len(s)
	if size == 0 {
		return 0
	}
	dict := make(map[byte]bool)
	maxSize := 0
	for i,j := 0, 0; i <= j && j < size; {
		flag, ok := dict[s[j]]
		if !flag || !ok {
			dict[s[j]] = true
			j++
		} else {
			dict[s[i]] = false
			i++
		}
		if j - i > maxSize {
			maxSize = j - i
		}
	}
	return maxSize
}
